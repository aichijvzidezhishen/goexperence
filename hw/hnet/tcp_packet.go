package hnet

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

// 包头长度：4字节（存储int32，最大支持4GB包体）
const HeaderLength = 4

// 默认超时时间
const DefaultTimeout = 30 * time.Second

// Packet 封装/解包工具类
type Packet struct {
	conn          net.Conn      // TCP连接
	timeout       time.Duration // 读写超时时间
	enableZip     bool          // 是否启用gzip压缩
	enableEncrypt bool          // 是否启用AES加密
	encryptKey    []byte        // AES加密密钥（32字节，AES-256）
}

// PacketOption 配置选项函数
type PacketOption func(*Packet)

// WithTimeout 设置超时时间
func WithTimeout(timeout time.Duration) PacketOption {
	return func(p *Packet) {
		p.timeout = timeout
	}
}

// WithZip 启用gzip压缩
func WithZip(enable bool) PacketOption {
	return func(p *Packet) {
		p.enableZip = enable
	}
}

// WithEncrypt 启用AES-256加密（密钥必须为32字节，建议用sha256生成）
func WithEncrypt(key string) PacketOption {
	return func(p *Packet) {
		// 用sha256处理密钥，确保是32字节
		hash := sha256.Sum256([]byte(key))
		p.encryptKey = hash[:]
		p.enableEncrypt = true
	}
}

// NewPacket 创建Packet实例
// options: 可选配置（超时、压缩、加密）
func NewPacket(conn net.Conn, options ...PacketOption) *Packet {
	p := &Packet{
		conn:    conn,
		timeout: DefaultTimeout, // 默认30秒超时
	}
	// 应用配置选项
	for _, opt := range options {
		opt(p)
	}
	return p
}

// -------------------------- 核心工具函数 --------------------------
// gzip压缩数据
func zipData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write(data)
	if err != nil {
		return nil, fmt.Errorf("压缩失败：%w", err)
	}
	if err := zw.Close(); err != nil {
		return nil, fmt.Errorf("压缩流关闭失败：%w", err)
	}
	return buf.Bytes(), nil
}

// gzip解压数据
func unzipData(data []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("解压失败：%w", err)
	}
	defer zr.Close()
	unzipData, err := io.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("读取解压数据失败：%w", err)
	}
	return unzipData, nil
}

// AES-256-CFB加密（带随机IV）
func encryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建加密器失败：%w", err)
	}

	// 生成随机IV（16字节，AES块大小）
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("生成IV失败：%w", err)
	}

	// CFB模式加密
	stream := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(data))
	stream.XORKeyStream(cipherText, data)

	// IV + 密文 返回（解密时需要IV）
	result := append(iv, cipherText...)
	return result, nil
}

// AES-256-CFB解密
func decryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建解密器失败：%w", err)
	}

	// 分离IV和密文
	if len(data) < aes.BlockSize {
		return nil, errors.New("密文长度不足")
	}
	iv := data[:aes.BlockSize]
	cipherText := data[aes.BlockSize:]

	// CFB模式解密
	stream := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

// -------------------------- 核心业务方法 --------------------------
// Pack 封包：压缩→加密→添加包头
func (p *Packet) Pack(data []byte) ([]byte, error) {
	var err error

	// 1. 压缩数据（如果启用）
	if p.enableZip {
		data, err = zipData(data)
		if err != nil {
			return nil, err
		}
	}

	// 2. 加密数据（如果启用）
	if p.enableEncrypt {
		data, err = encryptData(data, p.encryptKey)
		if err != nil {
			return nil, err
		}
	}

	// 3. 添加包头（包体长度）
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, int32(len(data)))
	if err != nil {
		return nil, fmt.Errorf("封包失败：%w", err)
	}
	_, err = buf.Write(data)
	if err != nil {
		return nil, fmt.Errorf("封包失败：%w", err)
	}

	return buf.Bytes(), nil
}

// Unpack 解包：读取包头→读取包体→解密→解压
func (p *Packet) Unpack() ([]byte, error) {
	// 设置读取超时
	p.conn.SetReadDeadline(time.Now().Add(p.timeout))

	// 1. 读取包头
	header := make([]byte, HeaderLength)
	n, err := p.conn.Read(header)
	if err != nil {
		return nil, fmt.Errorf("读取包头失败：%w", err)
	}
	if n != HeaderLength {
		return nil, fmt.Errorf("包头长度异常，预期%d字节，实际%d字节", HeaderLength, n)
	}

	// 2. 解析包体长度
	var bodyLen int32
	buf := bytes.NewReader(header)
	err = binary.Read(buf, binary.BigEndian, &bodyLen)
	if err != nil {
		return nil, fmt.Errorf("解析包头失败：%w", err)
	}
	if bodyLen <= 0 {
		return nil, fmt.Errorf("包体长度非法：%d", bodyLen)
	}

	// 3. 读取包体
	body := make([]byte, bodyLen)
	totalRead := 0
	for totalRead < int(bodyLen) {
		n, err := p.conn.Read(body[totalRead:])
		if err != nil {
			return nil, fmt.Errorf("读取包体失败：%w", err)
		}
		totalRead += n
	}

	// 4. 解密（如果启用）
	if p.enableEncrypt {
		body, err = decryptData(body, p.encryptKey)
		if err != nil {
			return nil, fmt.Errorf("解密失败：%w", err)
		}
	}

	// 5. 解压（如果启用）
	if p.enableZip {
		body, err = unzipData(body)
		if err != nil {
			return nil, fmt.Errorf("解压失败：%w", err)
		}
	}

	return body, nil
}

// Send 快速发送数据（带超时）
func (p *Packet) Send(data []byte) error {
	// 设置写入超时
	p.conn.SetWriteDeadline(time.Now().Add(p.timeout))

	packedData, err := p.Pack(data)
	if err != nil {
		return err
	}
	_, err = p.conn.Write(packedData)
	if err != nil {
		return fmt.Errorf("发送数据失败：%w", err)
	}
	return nil
}

// Recv 快速接收数据（带超时）
func (p *Packet) Recv() ([]byte, error) {
	return p.Unpack()
}

// Close 关闭连接
func (p *Packet) Close() error {
	return p.conn.Close()
}

// GetConn 获取原始连接（备用）
func (p *Packet) GetConn() net.Conn {
	return p.conn
}

// 生成加密密钥（辅助函数）：将任意字符串转为32字节的AES-256密钥
func GenerateAESKey(rawKey string) string {
	hash := sha256.Sum256([]byte(rawKey))
	return hex.EncodeToString(hash[:])
}

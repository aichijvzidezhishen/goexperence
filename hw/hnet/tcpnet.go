package hnet

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	MaxPacketSize = 1024
)

// TcpNet函数用于创建一个TCP服务器
func TcpNet() {
	// 监听8080端口
	listener, err := net.Listen("tcp", ":8080")
	// 如果监听失败，打印错误信息并返回
	if err != nil {
		fmt.Println("net listen err:", err.Error())
		return
	}
	// 关闭监听
	defer listener.Close()
	// 打印监听成功信息
	fmt.Println("Tcp server listening... : 8080")

	// 无限循环，等待客户端连接
	for {

		// 接受客户端
		conn, err := listener.Accept()
		// 如果接受失败，打印错误信息并继续循环
		if err != nil {
			fmt.Println("net accept err:", err.Error())
			continue
		}

		// 打印客户端连接信息
		fmt.Printf("new tcp connection from %s\n", conn.RemoteAddr())

		// 为每个客户端创建一个goroutine
		go handleTcpConnection(conn)

	}
}

func handleTcpConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, MaxPacketSize+4) // +4 bytes for the length prefix

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	// Read the length prefix
	length := binary.BigEndian.Uint32(buf[:4])
	if length > uint32(n-4) {
		fmt.Println("Invalid packet length", length, "n-4:", n-4)
		return
	}
	message := buf[4 : length+4]
	fmt.Printf("Tcp received %s,length:%d n-4 %d \n", message, length, n-4)
	conn.Write([]byte("ok"))
}

func StartUDPServer() {
	conn, err := net.ListenPacket("udp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// 关闭监听
	defer conn.Close()
	fmt.Println("UDP server listening... : 8081")

	buf := make([]byte, MaxPacketSize+4) // +4 bytes for the length prefix
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			continue
		}

		length := binary.BigEndian.Uint32(buf[:4])
		if length > uint32(n-4) {
			fmt.Println("Invalid packet length")
			continue
		}

		message := buf[4 : length+4]
		fmt.Printf("UDP received %s from %s\n", message, addr.String())
		conn.WriteTo([]byte("ok"), addr)
	}
}

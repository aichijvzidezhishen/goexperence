package http

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
	func (r *Request) AddCookie(c *Cookie)  // 添加 cookie

func (r *Request) BasicAuth() (username, password string, ok bool) // 返回请求基本认证中的用户名密码
func (r *Request) Cookie(name string) (*Cookie, error) // 返回指定名称的 Cookie
func (r *Request) Cookies() []*Cookie  // 返回所有 cookie
// 返回表单中指定 key 中的第一个文件对象
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
func (r *Request) FormValue(key string) string // 返回表单中 key 的值
func (r *Request) ParseForm() error  // 解析表单
func (r *Request) ParseMultipartForm(maxMemory int64) error // 解析带有文件的表单
func (r *Request) Referer() string  // 返回引用 URL
func (r *Request) SetBasicAuth(username, password string)  // 设置基本认证请求的用户名密码
func (r *Request) UserAgent() string // 返回请求的客户端代理
func (r *Request) Write(w io.Writer) error // 将请求写入文件
*/

func httpPostForm() {
	data := url.Values{
		"name": []string{"name"},
		"age":  []string{"20"},
		"addr": []string{"beijing", "shanghai", "guangzhou"},
	}
	res, err := http.PostForm("http://www.httpbin.org/post", data)
	if err != nil {
		fmt.Println("post from request failed", err)
		return
	}
	defer res.Body.Close()
	buf := make([]byte, 2048)
	n, err := res.Body.Read(buf)
	if err != nil {
		fmt.Println("read body failed", err)
	}
	fmt.Println(string(buf[:n]), "\n ---- body length---- :", n)
}

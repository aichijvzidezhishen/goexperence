package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jackdanger/collectlinks"
)

var visited = make(map[string]bool)

/*
func main() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("get error:", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	// ReadAll 读取 Body 中的所有数据，返回读取的数据和遇到的错误。
	// 如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
	// 所有数据，所以不会把 EOF 当做错误处理。
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println(string(body))
}
*/

func main() {
	url := "http://www.baidu.com/"
	queue := make(chan string)
	go func() {
		queue <- url
	}()
	for urli := range queue {
		download(urli, queue)
	}
	// url := "http://www.01happy.com/demo/accept.php"

}

func download(url string, queue chan string) {
	visited[url] = true
	timeout := time.Duration(5 * time.Second)

	client := &http.Client{
		Timeout: timeout,
	}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)
	for _, link := range links {
		absolute := urlJoin(link, url)
		if url != " " {
			if !visited[absolute] {
				fmt.Println("parse url", absolute)
				go func() {
					queue <- absolute
				}()
			}
		}
	}
}

func urlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()
}

// func download_post(url string) {
// 	//相对于download方法 抓取的页面完整
// 	//使用这个方法时，第二个参数要设置成"application/x-www-form-urlencoded",否侧参数无法传递
// 	resp, err := http.Post(url,
// 		"application/x-www-form-urlencoded",
// 		strings.NewReader("name=cjb"))
// 	if err != nil {
// 		fmt.Println("post error:", err)
// 	}

// 	defer resp.Body.Close() //关闭网页
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("read error:", err)
// 		return
// 	}

// 	fmt.Println(string(body))

// }

//使用复杂的请求时（需要在请求时设置头参，cookie之类的数据，方法封装成函数
// func download(url string) {
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", url, nil)

// 	//自定义header
// 	req.Header.Set("User-Agent", "Mozilla/4.0(compatibe;MSIE 6.0;Window NT 5.1)")

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("get error:", err)
// 		return
// 	}
// 	defer resp.Body.Close() //关闭网页

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		fmt.Println("read err:", err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }

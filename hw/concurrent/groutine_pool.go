package concurrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// http get 请求
func httpget(ch chan int) {
	resp, err := http.Get("http://localhost:8080/rest/api/users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	ch <- 1
}

func impMutiHttpGet() {
	start := time.Now()

	cslen := 100
	cs := make([]chan int, cslen)

	//
	for i := 0; i < cslen; i++ {
		cs[i] = make(chan int)
		go httpget(cs[i])
	}

	for _, v := range cs {
		<-v
	}
	since := time.Since(start)
	fmt.Println("process cost time ", since)

}

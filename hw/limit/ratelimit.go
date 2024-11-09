package limit

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

func RateLimit() {
	r := rate.Every(1 * time.Millisecond)
	limit := rate.NewLimiter(r, 10)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if limit.Allow() {
			fmt.Printf("success，当前时间：%s\n", time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("success，但是被限流了。。。\n")
		}
	})

	fmt.Println("http start ... ")
	_ = http.ListenAndServe(":8080", nil)
}

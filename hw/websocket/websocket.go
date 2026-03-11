package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 定义 WebSocket 连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有来源的连接
		return true
	},
}

func implWebSocket() {
	http.HandleFunc("/ws", handleWebSocket)

	// 启动HTTP服务器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

// 实现 WebSocket 连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将 http 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// 处理错误
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	fmt.Println("client connected ")

	// 处理 WebSocket 连接
	for {
		// 读取消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			// 处理错误
			log.Printf("Error reading message: %v", err)
			break
		}
		//打印收到的消息
		fmt.Printf("Received message: %s\n", msg)
		// 处理消息
		// ...

		// 发送消息
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			// 处理错误
			log.Printf("Error writing message: %v", err)
			break
		}
	}
	fmt.Println("client disconnected ")
}

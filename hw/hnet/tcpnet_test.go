package hnet

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTcpNet(t *testing.T) {
	// Test case 1: Test the TCP server's ability to listen on a specific port
	go TcpNet()
	time.Sleep(1 * time.Second) // Wait for the server to start

	// Test case 2: Test the TCP server's ability to accept client connections
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Errorf("Failed to connect to TCP server: %v", err)
	}
	msg := "Hello, server!"
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(msg)))
	_, err = conn.Write(length)
	if err != nil {
		t.Errorf("Failed to send data to TCP server: %v", err)
		return
	}
	_, err = conn.Write([]byte(msg))
	if err != nil {
		t.Errorf("Failed to send data to TCP server: %v", err)
		return
	}
	fmt.Println("Connected to TCP server")
	defer conn.Close()

	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		conn, err := net.Dial("tcp", "localhost:8080")
	// 		if err != nil {
	// 			t.Errorf("Failed to connect to TCP server: %v", err)
	// 		}
	// 		defer conn.Close()
	// 	}()
	// }

	// fmt.Println("Connected to TCP server multiple times succ")
}

func TestStartUDPServer(t *testing.T) {
	// tests := []struct {
	// 	name string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		StartUDPServer()
	// 	})
	// }
	StartUDPServer()
	time.Sleep(10 * time.Second) // Wait for the server to start
}

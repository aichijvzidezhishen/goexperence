package hnet

import (
	"fmt"
	"testing"
)

func TestGetLocalIp(t *testing.T) {
	ip, err := GetLocalIp()
	if err != nil {
		return
	}
	fmt.Println("local ip: ", ip)
}

//

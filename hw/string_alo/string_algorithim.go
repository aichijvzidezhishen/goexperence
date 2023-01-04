package string_alo

import (
	"fmt"
	"strconv"
	"strings"
)

func PointMove(input string) (int, int) {
	x, y := 0, 0

	opts := strings.Split(input, ";")

	for i := 0; i < len(opts); i++ {
		opt := opts[i]
		optBytes := []byte(opt)

		if len(optBytes) < 1 {
			continue
		}
		fmt.Println("optBytes", string(optBytes[1:]))
		distanceStr := string(optBytes[1:])
		distance, err := strconv.ParseUint(distanceStr, 0, 64)
		if err != nil {
			continue
		}
		switch optBytes[0] {
		case 'A':
			x -= int(distance)

		case 'D':
			x += int(distance)

		case 'W':
			y += int(distance)

		case 'S':
			y -= int(distance)
		}
	}

	return x, y
}

// 识别有效的IP地址和掩码并进行分类统计
func IdentifyIp(ipstr string) {
	iplst := strings.Split(ipstr, "~")
	// fmt.Print")
	if strings.HasPrefix(iplst[0], "0") {
		fmt.Println("illegal input ---- ")
	}
}
func CheckIp(ip []int) (byte, bool) {
	if !IsOk(ip) {
		return 'f', false
	}

	if ip[0] >= 0 && ip[0] <= 126 {
		return 'a', ip[0] == 10
	}

	if ip[0] >= 128 && ip[0] <= 191 {
		return 'b', ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31
	}

	if ip[0] >= 192 && ip[0] <= 223 {
		return 'c', ip[0] == 192 && ip[1] == 168
	}

	if ip[0] >= 224 && ip[0] <= 239 {
		return 'd', false
	}
	if ip[0] >= 240 && ip[0] <= 255 {
		return 'e', false
	}

	return 'n', false
}

func CheckMask(ip []int) bool {
	if ip == nil {
		return false
	}
	var res uint32
	for i := 0; i < len(ip); i++ {
		res = res<<8 + uint32(ip[i])
	}

	if res == 0 || res == 0xffffffff {
		return false
	}

	const high = 0x80000000
	for {
		if (res & high) == 0 {
			break
		}
		res = res << 1
	}
	return res == 0
}

func IsOk(ip []int) bool {
	if ip == nil {
		return false
	}
	res := true
	for i := 0; i < len(ip); i++ {
		res = res && ip[1] >= 0 && ip[i] <= 255
	}

	return res
}

func parse(s string) []int {
	sp := strings.Split(s, ".")
	if len(sp) != 4 {
		return nil
	}

	res := []int{}
	for i := 0; i < 4; i++ {
		val, err := strconv.Atoi(sp[i])
		if err != nil {
			return nil
		}
		res = append(res, val)
	}
	return res
}

//

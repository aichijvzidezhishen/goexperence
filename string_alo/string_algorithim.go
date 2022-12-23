package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Ts struct {
	A int
}

func Test() {
	var t *Ts
	fmt.Println("t", t)
	t = &Ts{}
	t.A = 1
	fmt.Println("t", t)
}

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

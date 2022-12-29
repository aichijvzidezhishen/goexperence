package main

import (
	"fmt"
	"holiday/day04/encapsulation/model"
)

func main() {
	p := model.NewPerson("jack")
	p.Setage(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, "age=", p.GetAge(), "sal=", p.GetSal())
}

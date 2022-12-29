package main

import "fmt"

/*
	定义动物接口：生，死，活着
	定义动物实现类：鸟，鱼，野兽（跑，捕食）
	继承野兽：实现老虎，实现人
	业务场景：工作日动物都活着，周末人出来捕食，死了
	其他动物死了，人死了

*/

type animal interface {
	chusheng()
	godie()
	live()
}
type bird struct {
	name string
}

func (b *bird) chusheng() {
	fmt.Println("鸟破蛋")
}
func (b *bird) godie() {
	fmt.Println("鸟死了")
}
func (b *bird) live() {
	fmt.Println("鸟在唱歌")
}

type fish struct {
	name string
}

func (f *fish) chusheng() {
	fmt.Println("鱼破蛋")
}
func (f *fish) godie() {
	fmt.Println("鱼死了")
}
func (f *fish) live() {
	fmt.Println("鱼在唱歌")
}

type beast struct {
	name string
}

func (be *beast) chusheng() {
	fmt.Println("野兽破蛋")
}
func (be *beast) godie() {
	fmt.Println("野兽死了")
}
func (be *beast) live() {
	fmt.Println("野兽在唱歌")
}
func (be *beast) run() {
	fmt.Println("野兽在跑")
}
func (be *beast) eat(ani animal) {
	fmt.Println("野兽捕食")
}

type tiger struct {
	beast
}

type human struct {
	beast
}

func main() {
	animals := make([]animal, 0)
	animals = append(animals, &bird{})
	animals = append(animals, &fish{})
	animals = append(animals, &tiger{beast{name: "啊荣"}})
	h := &human{beast{name: "若"}}
	animals = append(animals, h)

	for _, v := range animals {
		switch v.(type) {
		case *bird:
			fmt.Println("ok")
		}
	}
}

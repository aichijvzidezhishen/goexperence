package dfa

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
	// "github.com/bean-du/dfa"
)

func TestCheck(t *testing.T) {
	list := []string{"国", "发票", "在论发", "国家"}

	fmt.Println("sss", list)
	fda := New()
	fda.AddBadWords(list)
	// succDist, sfound, starget := fda.Check("中发票")
	// fmt.Println("succDist", succDist, "found", sfound, "target", starget)

	// failDist, ffound, ftarget := fda.Check("/da发国")
	// fmt.Println("failDist", failDist, "ffound", ffound, "ftarget", ftarget)
	failDist1, ffound1, ftarget1 := fda.Check("/发国在 国家")
	fmt.Println("failDist--", failDist1, "ffound", ffound1, "ftarget", ftarget1)
}

func BenchmarkCheck(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	list := []string{"中国", "发票", "发论"}
	file, err := os.Open("./censor.txt")
	if err != nil {
		fmt.Println("err")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		a, _, c := reader.ReadLine()
		if c == io.EOF {
			break
		}
		list = append(list, string(a))
	}
	fda := New()
	fda.AddBadWords(list)

	for i := 0; i < b.N; i++ {
		fda.Check("dasd中国zgss人 法轮功")
	}
}

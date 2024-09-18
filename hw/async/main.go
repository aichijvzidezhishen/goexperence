package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gendataII(numGoRoutines, numElements int) []int {
	var wg sync.WaitGroup
	wg.Add(numGoRoutines)
	rand.Seed(time.Now().UnixNano())
	result := make([]int, numElements)

	ssi := make([][]int, numGoRoutines)
	partitionElements := numElements / numGoRoutines
	for i := 0; i < numGoRoutines; i++ {
		ssi[i] = make([]int, 0, partitionElements)
		go func(i int) {
			defer wg.Done()
			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s)
			offset := i * partitionElements
			for j := 0; j < partitionElements; j++ {
				result[j+offset] = r.Intn(10)
			}
		}(i)
	}
	wg.Wait()

	// Create the result by appending all the sub-slices into one result slice
	// var res []int
	// for i := 0; i < numGoRoutines; i++ {
	// 	res = append(res, ssi[i]...)
	// }

	return result
}
func gendata(numGoRoutines, numElements int) []int {
	var wg sync.WaitGroup
	wg.Add(numGoRoutines)

	res := make([]int, numElements)                  // result slice
	partitionElements := numElements / numGoRoutines // number of elements per goroutine

	for i := 0; i < numGoRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			offset := i * partitionElements // offset to calculate the index for the data insertion
			for j := 0; j < partitionElements; j++ {
				res[j+offset] = r.Intn(10)
			}
		}(i)
	}
	wg.Wait()

	return res
}
func main() {
	n := 100000000

	// One Go routine
	g := 1
	t1 := time.Now()
	res1 := gendataII(g, n)

	t2 := time.Now()
	fmt.Println("Data generation of", n, "elements with", g, "Go routines took", t2.Sub(t1), "produce ", len(res1))

	// Two Go routines
	g = 2
	t1 = time.Now()
	res2 := gendataII(g, n)
	t2 = time.Now()
	fmt.Println("Data generation of", n, "elements with", g, "Go routines took", t2.Sub(t1), "produce ", len(res2))
}

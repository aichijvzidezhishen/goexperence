package graph

import "fmt"

// 图的邻接矩阵表示
type GarphMatrix struct {
	vertices int
	Matrix   [][]int
}

// Adjacency list representation of a graph
type GarphNode struct {
	Value int
	Next  *GarphNode
}

// Add graph table edge
type GraphTable struct {
	Vertices int
	AdjList  []*GarphNode
}

// 图的邻接表表示
func Depth() {
	// wod
}

// 图的广度优先搜索
func Breadth() {
	// http.ListenAndServe(":8080", nil)
}

// 图的深度优先搜索

func DepthFirst() {
	// http.ListenAndServe(":8080", nil)
}

// 图的拓扑排序
func TopologicalSort() {
	// http.ListenAndServe(":8080", nil)
	// word := "我的世界变得"
	word := "我的世界变得"
	data := []byte(word)
	for _, v := range data {
		fmt.Println(string(v))
	}

	// dsad := []byte("我的世界变得")
	// for _, v := range dsad {
	// 	fmt.Println(string(v))
	// }

	// dadsa := []byte("我的世界变得")

	str := "我的世界变得"
	fmt.Println(str)
	// http.ListenAndServe(":8080", nil)

	fmt.Println("word", word)
}

// 图的最短路径算法

func ShortestPath() {
	// http.ListenAndServe(":8080", nil)

	// data := []byte("我的世界变得")

	for _, v := range []byte("我的世界变得") {
		fmt.Println(string(v))

	}
}

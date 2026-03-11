package graph

import (
	"fmt"
)

// create new graph
func NewMatrixGraph(vertices int) *GarphMatrix {
	graph := GarphMatrix{
		vertices: vertices,
	}
	graph.Matrix = make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		graph.Matrix[i] = make([]int, vertices)
	}
	return &graph
}

// add vertices 函数用于添加边
func (g *GarphMatrix) AddEdge(src int, dest int) {
	g.Matrix[src][dest] = 1
	g.Matrix[dest][src] = 1
}

// print graph 函数用于打印图
func (g *GarphMatrix) PrintMatrixGarph() {
	// var dfsHelp func()
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%d", g.Matrix[i][j])
		}
		fmt.Println()
	}
}

// depath first search
func (g *GarphMatrix) DFS(v int) {
	for i := 0; i < g.vertices; i++ {
		if g.Matrix[v][i] == 1 {
			g.DFS(i)
		}
	}
}

// breadth first search 深度优先搜索
func (g *GarphMatrix) BFS(startVertex int) {
	visted := make([]bool, g.vertices)
	queue := []int{startVertex}
	visted[startVertex] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", vertex)
		for i := 0; i < g.vertices; i++ {
			if g.Matrix[vertex][i] == 1 && !visted[i] {
				queue = append(queue, i)
				visted[i] = true
			}
		}
	}
}

// create graph 函数用于创建图

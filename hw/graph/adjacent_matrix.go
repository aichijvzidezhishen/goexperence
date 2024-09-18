package graph

import "fmt"

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

// add vertices
func (g *GarphMatrix) AddEdge(u int, v int) {
	g.Matrix[u][v] = 1
	g.Matrix[v][u] = 1
}

// print graph
func (g *GarphMatrix) PrintMatrixGarph() {
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%d", g.Matrix[i][j])
		}
		fmt.Println()
	}
}

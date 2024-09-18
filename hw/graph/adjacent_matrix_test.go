package graph

import "testing"

func TestGarphMatrix_PrintMatrixGarph(t *testing.T) {
	graph := NewMatrixGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	graph.PrintMatrixGarph()
}

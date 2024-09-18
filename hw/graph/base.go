package graph

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

//

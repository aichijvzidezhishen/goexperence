package graph

func NewAdjacentTableGraph(vertices int) *GraphTable {
	return &GraphTable{
		Vertices: vertices,
		AdjList:  make([]*GarphNode, vertices),
	}
}

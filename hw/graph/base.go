package graph

import (
	"errors"
	"fmt"
	"strings"
)

// 无向无权图（邻接表实现）
type UndirectedGraph struct {
	adj map[string][]string // 邻接表：key=顶点，value=相邻顶点列表
}

func NewUndirectGraph() *UndirectedGraph {
	return &UndirectedGraph{
		adj: make(map[string][]string),
	}
}

func (c *UndirectedGraph) AddVertex(v string) {
	if _, exist := c.adj[v]; !exist {
		c.adj[v] = []string{}
	}
}

func (c *UndirectedGraph) AddEdge(e1, e2 string) error {
	// 检查顶点
	if _, exist := c.adj[e1]; !exist {
		return errors.New("顶点" + e1 + "不存在")
	}

	if _, exist := c.adj[e2]; !exist {
		return errors.New("顶点" + e1 + "不存在")
	}

	// 避免重复添加
	if !c.HasEdge(e1, e2) {
		c.adj[e1] = append(c.adj[e1], e2)
	}

	if c.HasEdge(e2, e1) {
		c.adj[e2] = append(c.adj[e2], e1)
	}
	return nil
}

func (c *UndirectedGraph) HasEdge(e1, e2 string) bool {
	for _, v := range c.adj[e1] {
		if v == e2 {
			return true
		}
	}
	return false
}

// 删除边
func (c *UndirectedGraph) RemoveEdge(e1, e2 string) bool {
	// c.adj[e1][e2]
	adj1, hasV1 := c.adj[e1]
	adj2, hasV2 := c.adj[e2]
	if !hasV1 || !hasV2 {
		return false
	}

	hasEdge1 := sliceContains(adj1, e2)
	hasEdge2 := sliceContains(adj1, e1)
	if !hasEdge2 || !hasEdge1 {
		return false
	}

	// 从e1中删除
	c.adj[e1] = removeFromSlice(adj1, e2)
	// 从e2中删除
	c.adj[e2] = removeFromSlice(adj2, e1)
	return true

}

func sliceContains(list []string, ele string) bool {
	for _, v := range list {
		if v == ele {
			return true
		}

	}
	return false
}

func removeFromSlice(list []string, target string) []string {
	for i, v := range list {
		if v == target {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func (c *UndirectedGraph) Print() string {
	var sb strings.Builder
	sb.WriteString("无向无权图：\n")
	for i, neighbors := range c.adj {
		sb.WriteString(fmt.Sprintf("%s :%v \b", i, neighbors))
	}
	return sb.String()
}

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
}

// 图的广度优先搜索
// Breadth 函数定义，可能用于实现广度优先搜索或相关功能
// 目前该函数被注释掉，可能处于开发调试阶段
func Breadth() {
	// http.ListenAndServe(":8080", nil)  // 这行代码被注释掉，原本用于启动一个监听在8080端口的HTTP服务器
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

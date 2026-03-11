package tree

import (
	"fmt"
	"strconv"
)

type BtreeNode struct {
	keys   []int
	values map[int]interface{}
	childs []*BtreeNode
}

type Node struct {
	value int
	left  *Node
	right *Node
}

// Print 是一个公共方法，用于启动打印过程
func (n *Node) Print() {
	if n == nil {
		fmt.Println("(Tree is empty)")
		return
	}
	fmt.Println("---------------------------------")
	printTree(n, 0, false) // 根节点没有父节点，isLeft值不重要
	fmt.Println("---------------------------------")
}

// printTree 是一个内部递归函数，负责实际的打印工作
// depth 用来控制缩进
// isLeft 用来判断当前节点是其父节点的左子节点还是右子节点，以绘制不同的连接符
func printTree(node *Node, depth int, isLeft bool) {
	if node == nil {
		return
	}

	// 1. 递归打印右子树
	printTree(node.right, depth+1, false)

	// 2. 打印当前节点
	for i := 0; i < depth; i++ {
		fmt.Print("    ")
	}

	if depth > 0 {
		if isLeft {
			fmt.Print("└── ")
		} else {
			fmt.Print("┌── ")
		}
	}

	fmt.Println(strconv.Itoa(node.value))

	// 3. 递归打印左子树
	printTree(node.left, depth+1, true)
}

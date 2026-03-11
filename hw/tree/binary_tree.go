package tree

func NewTreeNode(val int) *Node {
	return &Node{
		value: val,
		left:  nil,
		right: nil,
	}
}

// 深度优先
func DFSDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return maxInt(DFSDepth(root.left), DFSDepth(root.right)) + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 广度优先
func BFSDepth(root *Node) int {
	if root == nil {
		return 0
	}

	queue := []*Node{root}
	depth := 0 // 深度初始化为0
	// 深度初始化为0
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // 取出头节点
			queue = queue[1:] // 出队

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		depth++
	}
	return depth
}

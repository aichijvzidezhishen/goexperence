package tree

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (b *BST) Insert(value int) {
	nNode := &Node{value, nil, nil}

	if b.root == nil {
		b.root = nNode
	} else {
		insertNode(b.root, nNode)
	}

}

func insertNode(root, inNode *Node) {
	if inNode.value < root.value { // 放在根结点的左边
		if root.left == nil {
			root.left = inNode
		} else {
			insertNode(root.left, inNode)
		}
	} else if inNode.value > root.value {
		if root.right == nil {
			root.right = inNode
		} else {
			insertNode(root.right, inNode)
		}
	}
}

// 删除一个元素
func remove(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	// 从左边找
	if value < root.value {
		root.right, existed = remove(root.right, value)
		return root, existed
	}

	//从右边找
	if value < root.value {
		root.left, existed = remove(root.left, value)
		return root, existed
	}

	// 如果此节点是正要删除的节点，那么就返回此节点，同时返回之前可能需要调整
	existed = true
	//
	if root.left == nil && root.right == nil {
		root = nil
		return root, existed
	}
	//
	if root.left == nil {
		root = root.right
		return root, existed
	}
	if root.right == nil {
		root = root.left
		return root, existed
	}

	smallestInRight, _ := min(root.right)

	//提升
	root.value = smallestInRight
	// 从右边子树中移除此节点
	root.right, _ = remove(root.right, smallestInRight)
	return nil, false
}

// Min
func (b *BST) Min() (int, bool) {
	return min(b.root)
}

func min(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node

	// 从左边找
	for {
		if n.left == nil {
			return int(n.value), true
		}
		n = n.left
	}
}

func (b *BST) Max() (int, bool) {
	return max(b.root)
}

func max(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node

	for {
		if n.right == nil {
			return int(n.value), true
		}
		n = n.right
	}
}

package tree

// BST

type BST struct {
	root *Node
}

// func (b *BST) Insert(value int) {
// 	nNode := &Node{value, nil, nil}

// 	if b.root == nil {
// 		b.root = nNode
// 	} else {
// 		insertNode(b.root, nNode)
// 	}

// }

// insertNode 函数用于将新节点插入到二叉搜索树中
// 参数:
//
//	root - 当前子树的根节点
//	in - 要插入的新节点
//
// API 设计不够灵活：这个函数要求调用者必须先创建一个 Node 对象 (in) 再传入。一个更常见、更灵活的 API 是直接传入一个值（如 int），由函数内部负责创建节点。这使得调用代码更简洁。
// func insertNode(root, in *Node) {
// 	// 如果新节点的值小于当前根节点的值，则向左子树插入
// 	if in.value < root.value {
// 		// 如果左子树为空，则直接将新节点作为左子树
// 		if root.left == nil {
// 			root.left = in
// 		} else {
// 			// 否则递归地向左子树插入
// 			insertNode(root.left, in)
// 		}
// 	} else if in.value > root.value {
// 		// 如果新节点的值大于当前根节点的值，则向右子树插入
// 		// 如果右子树为空，则直接将新节点作为右子树
// 		if root.right == nil {
// 			root.right = in
// 		} else {
// 			// 否则递归地向右子树插入
// 			insertNode(root.right, in)
// 		}
// 	}
// 	// 如果新节点的值等于当前根节点的值，则不进行任何操作（二叉搜索树不允许重复值）
// }

func (c *Node) OptInsertNode(val int) *Node {
	if c == nil {
		return &Node{value: val}
	}

	if val < c.value {
		c.left = c.left.OptInsertNode(val)
	} else if val > c.value {
		c.right = c.right.OptInsertNode(val)
	}
	return c
}

// 删除一个元素
func remove(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	// 从左边找
	if value < root.value {
		root.left, existed = remove(root.left, value)
		return root, existed
	}

	//从右边找
	if value > root.value {
		root.right, existed = remove(root.right, value)
		return root, existed
	}

	// 如果此节点是正要删除的节点，那么就返回此节点，同时返回之前可能需要调整
	existed = true

	//
	if root.left == nil && root.right == nil {
		root = nil
		return nil, existed
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
	return root, existed
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

func bstToSlice(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, bstToSlice(root.left)...)
	result = append(result, root.value)
	result = append(result, bstToSlice(root.right)...)
	return result
}

func InsertSliceToBST(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	root := &Node{value: nums[0]}
	for _, v := range nums[1:] {
		// insertNode(root, &Node{v, nil, nil})
		root.OptInsertNode(v)
	}
	return root
}

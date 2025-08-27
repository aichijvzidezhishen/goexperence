package tree

type BPlusTree struct {
	root *Node
}

type BtreeNode struct {
	keys   []int
	values map[int]interface{}
	childs []*BtreeNode
}

func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		root: &BtreeNode{
			keys:   make([]int, 0),
			values: make(map[int]interface{}),
			childs: make([]*BtreeNode, 0),
		},
	}
}

func (t *BPlusTree) Insert(key int, value interface{}) {
	node := t.root
	for {
		// 如果当前节点是叶子节点，则直接插入
		if len(node.childs) == 0 {
			node.keys = append(node.keys, key)
			node.values[key] = value
			break
		}

		// 找到要插入的子节点
		var child *BtreeNode
		for i := 0; i < len(node.keys); i++ {
			if key < node.keys[i] {
				child = node.childs[i]
				break
			}
		}
		if child == nil {
			child = node.childs[len(node.keys)]
		}

		// 进入子节点
		node = child
	}
}

func (t *BPlusTree) Search(key int) interface{} {
	node := t.root
	for {
		// 如果当前节点是叶子节点，则查找键对应的值
		if len(node.childs) == 0 {
			return node.values[key]
		}

		// 找到要查找的子节点
		var child *BtreeNode
		for i := 0; i < len(node.keys); i++ {
			if key < node.keys[i] {
				child = node.childs[i]
				break
			}
		}

		//
		if child == nil {
			child = node.childs[len(node.keys)]
		}

		// 进入子节点
		node = child
	}
}

// func main() {
// 	tree := NewBPlusTree()

// 	tree.Insert(1, "value1")
// 	tree.Insert(2, "value2")
// 	tree.Insert(3, "value3")
// 	tree.Insert(4, "value4")
// 	tree.Insert(5, "value5")

// 	fmt.Println(tree.Search(3)) // 输出: value
// 	fmt.Println(tree.Search(6)) // 输出: nil
// }

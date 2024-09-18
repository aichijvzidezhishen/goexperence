package tree

// 定义B+ 树节点
type BPlusTreeNode struct {
	Keys      []int //
	IsLeaf    bool
	ChildNode []*BPlusTreeNode
	Next      *BPlusTreeNode // 叶子结点的下一个结点
}

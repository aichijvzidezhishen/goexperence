package tree

import (
	"log"
	"testing"
)

func TestRbtree_Insert(t *testing.T) {
	rbtree := New()

	rbtree.Insert(&RedNode{rbtree.NIL, rbtree.NIL, rbtree.NIL, RED, Int(10)})
	rbtree.Insert(&RedNode{rbtree.NIL, rbtree.NIL, rbtree.NIL, RED, Int(9)})
	rbtree.Insert(&RedNode{rbtree.NIL, rbtree.NIL, rbtree.NIL, RED, Int(8)})
	rbtree.Insert(&RedNode{rbtree.NIL, rbtree.NIL, rbtree.NIL, RED, Int(6)})
	rbtree.Insert(&RedNode{rbtree.NIL, rbtree.NIL, rbtree.NIL, RED, Int(7)})

	log.Println("rbtree counts : ", rbtree.count)

	log.Println("------ ", rbtree.root.Item)
	log.Println("----", rbtree.root.Left.Item, "---", rbtree.root.Right.Item)
	log.Println("--", rbtree.root.Left.Left.Item, "-", rbtree.root.Left.Right.Item)
}

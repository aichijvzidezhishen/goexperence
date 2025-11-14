package tree

import (
	"fmt"
	"testing"
)

func TestBFSDepth(t *testing.T) {
	// ex1 ： 空树
	var root1 *Node
	fmt.Printf("test case1 ", BFSDepth(root1))

	root2 := &Node{value: 1}

}

package tree

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		tree      []int // 初始树的切片表示
		value     int   // 要删除的值
		wantTree  []int // 期望的树的切片表示
		wantExist bool  // 期望是否存在该值
	}{
		// {
		// 	name:      "Delete from empty tree",
		// 	tree:      []int{},
		// 	value:     5,
		// 	wantTree:  []int{},
		// 	wantExist: false,
		// },
		{
			name:      "Delete non-existent value",
			tree:      []int{5, 3, 7, 2, 4, 6, 8},
			value:     4,
			wantTree:  []int{2, 3, 5, 6, 7, 8},
			wantExist: true,
		},
		{
			name:      "Delete leaf node (2)",
			tree:      []int{5, 3, 7, 2, 4, 6, 8},
			value:     11,
			wantTree:  []int{2, 3, 4, 5, 6, 7, 8},
			wantExist: false,
		},
		// {
		// 	name:      "Delete node with only right child (7)",
		// 	tree:      []int{5, 3, 8, 2, 4, 7}, // 7 只有右子节点 8
		// 	value:     7,
		// 	wantTree:  []int{2, 3, 4, 5, 8},
		// 	wantExist: true,
		// },
		// {
		// 	name: "Delete node with only left child (4)",
		// 	tree: []int{5, 3, 7, 2, 4, 6, 8}, // 4 是叶子节点，我们改一下树结构
		// 	// tree: []int{5, 3, 7, 2, 4, 6, 8}, // 为了测试，我们假设 4 有一个左子节点 3.5 (但我们的insert函数不支持)
		// 	// 更好的测试方式是手动构建树
		// 	value:     4,                       // 假设 4 是只有左子节点的节点
		// 	wantTree:  []int{2, 3, 5, 6, 7, 8}, // 假设 4 被其左子节点替换
		// 	wantExist: true,
		// },
		// {
		// 	name:      "Delete node with two children (3)",
		// 	tree:      []int{5, 3, 7, 2, 4, 6, 8},
		// 	value:     3,
		// 	wantTree:  []int{2, 4, 5, 6, 7, 8}, // 3 被右子树最小值 4 替换
		// 	wantExist: true,
		// },
		// {
		// 	name:      "Delete root node (5)",
		// 	tree:      []int{5, 3, 7, 2, 4, 6, 8},
		// 	value:     5,
		// 	wantTree:  []int{2, 3, 4, 6, 7, 8}, // 5 被右子树最小值 6 替换
		// 	wantExist: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := InsertSliceToBST(tt.tree)
			root.Print()
			newRoot, existed := remove(root, tt.value)

			// 检查返回的存在性标志
			if existed != tt.wantExist {
				t.Errorf("remove() existed = %v, want %v", existed, tt.wantExist)
			}

			// 检查树的结构是否正确
			gotTree := bstToSlice(newRoot)
			if !reflect.DeepEqual(gotTree, tt.wantTree) {
				t.Errorf("remove() tree = %v, want %v", gotTree, tt.wantTree)
			}
			root.Print()
		})
	}
}

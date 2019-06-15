package main

import (
	"fmt"

	"github.com/nurudianto/data-structure/btree"
	"github.com/nurudianto/data-structure/stack"
)

func main() {
	//build tree
	datas := []int{9, 2, 10, 1, 3}
	tree := btree.NewBinaryTree()
	for _, v := range datas {
		tree.Add(v)
	}
	fmt.Println(isValidBST(tree.Root))
}

// Valid BST will have ascending order if we iterate using indorder traversal
func isValidBST(root *btree.Node) bool {
	if root == nil {
		return true
	}

	s := stack.NewStack()
	var prev *btree.Node
	for root != nil || s.Length() > 0 {
		tmp := root
		for tmp != nil {
			s.Push(tmp)
			tmp = tmp.Left
		}

		root = s.Pop().(*btree.Node)
		if prev != nil && prev.Data >= root.Data {
			return false
		}
		prev = root
		root = root.Right
	}
	return true
}

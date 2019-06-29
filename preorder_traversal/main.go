package main

import (
	"fmt"

	"github.com/nurudianto/data-structure/stack"

	"github.com/nurudianto/data-structure/btree"
)

func main() {
	//build tree
	datas := []int{9, 2, 10, 1, 3}
	tree := btree.NewBinaryTree()
	for _, v := range datas {
		tree.Add(v)
	}
	fmt.Println(preorderTraversal(tree.Root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *btree.Node) []int {
	var results []int

	s := stack.NewStack()
	s.Push(root)
	for s.Length() > 0 {
		data := s.Pop().(*btree.Node)
		results = append(results, data.Data)

		if data.Right != nil {
			s.Push(data.Right)
		}

		if data.Left != nil {
			s.Push(data.Left)
		}
	}
	return results
}

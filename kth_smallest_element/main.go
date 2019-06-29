package main

import (
	"github.com/nurudianto/data-structure/btree"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *btree.Node, k int) int {
	if root == nil {
		return 0
	}
	s := newStack()

	for root != nil || len(s.datas) > 0 {
		for root != nil {
			s.push(root)
			root = root.Left
		}

		root = s.pop().(*btree.Node)
		if k == 1 {
			return root.Data
		}
		k--
		root = root.Right
	}
	return 0
}

type stack struct {
	datas []interface{}
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(data interface{}) {
	s.datas = append(s.datas, data)
}

func (s *stack) pop() interface{} {
	leng := len(s.datas)
	d := s.datas[leng-1]
	s.datas = s.datas[:leng-1]
	return d
}

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
	fmt.Println(inorderTraversal(tree.Root))
}

// Inorder Traversal Binary Tree:
// 1. Loop left node of the tree until the ends.
// 2. Store the node into the stack.
// 3. If the end of leftside reach, get node from the top of the stack and get its right child.
// 4. Do 1-3 steps repeatly until the stack is empty
func inorderTraversal(root *btree.Node) []int {
	res := make([]int, 0)
	stack := stack.NewStack()
	curr := root
	for curr != nil || stack.Length() != 0 {
		for curr != nil {
			res = append(res, curr.Data)
			stack.Push(curr)
			curr = curr.Left
		}

		curr = stack.Pop().(*btree.Node)
		curr = curr.Right

	}
	return res
}

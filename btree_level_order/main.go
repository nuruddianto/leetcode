package main

import (
	"fmt"

	"github.com/nurudianto/leetcode/data_structure/btree"
)

func main() {
	data := []int{3, 9, 20, 0, 0, 15, 7}
	tree := btree.NewTreeNode(data)
	fmt.Println(levelOrder(tree.Root))
}

func levelOrder(root *btree.Node) [][]int {
	marker := &btree.Node{Val: -1}

	queue := make([]*btree.Node, 0)
	queue = append(queue, root, marker)
	tmp := make([]int, 0)
	result := make([][]int, 0)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == nil || current.Val == 0 {
			continue
		}
		if current.Val == -1 {
			if len(tmp) > 0 {
				result = append(result, tmp)
			}
			tmp = make([]int, 0)
			if len(queue) > 0 {
				queue = append(queue, marker)
			}
		} else {
			tmp = append(tmp, current.Val.(int))
			queue = append(queue, current.Left, current.Right)
		}

	}

	return result
}

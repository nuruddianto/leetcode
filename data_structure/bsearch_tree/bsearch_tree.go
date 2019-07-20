package bsearch_tree

import (
	"errors"
	"fmt"
)

var ErrorNullInstace = errors.New("Binary Tree has not been initialized")

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type TreeNode struct {
	Root *Node
}

func NewBTree() *TreeNode {
	return &TreeNode{
		Root: &Node{},
	}
}

func (b *TreeNode) Add(data int) {
	if b.Root == nil {
		fmt.Println(ErrorNullInstace.Error())
		return
	}

	newData := &Node{Data: data}
	addRecursively(b.Root, newData)
}

func addRecursively(curr *Node, data *Node) *Node {
	if curr == nil {
		return data
	}

	if curr.Data == 0 {
		curr.Data = data.Data
		return data
	}

	if data.Data < curr.Data {
		curr.Left = addRecursively(curr.Left, data)
	} else if data.Data > curr.Data {
		curr.Right = addRecursively(curr.Right, data)
	} else {
		return curr
	}

	return curr
}

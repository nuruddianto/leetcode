package btree

type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

type TreeNode struct {
	Root *Node
}

func NewTreeNode(data []int) *TreeNode {
	root := &Node{Val: data[0]}
	setTree(0, data, root)
	return &TreeNode{
		Root: root,
	}
}

func setTree(idx int, data []int, root *Node) {
	if idx >= len(data) {
		return
	}

	left := 2*idx + 1
	rigth := 2*idx + 2

	if left < len(data) {
		root.Left = &Node{Val: data[left]}
		setTree(left, data, root.Left)
	}

	if rigth < len(data) {
		root.Right = &Node{Val: data[rigth]}
		setTree(rigth, data, root.Right)
	}
}

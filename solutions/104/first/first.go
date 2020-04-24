package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
	}

	root.Left = &TreeNode{
		Val: 2,
	}

	root.Right = &TreeNode{
		Val: 3,
	}

	root.Right.Right = &TreeNode{
		Val: 4,
	}

	root.Right.Right.Right = &TreeNode{
		Val: 5,
	}

	fmt.Println(maxDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var re int
	if root == nil {
		return re
	}

	re++

	re += better(maxDepth(root.Left), maxDepth(root.Right))

	return re
}

func better(i, j int) int {
	if i > j {
		return i
	}

	return j
}

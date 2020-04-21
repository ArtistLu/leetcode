package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := len(stack) - 1
		cur := stack[top]
		res = append(res, cur.Val)
		root = cur.Right
		stack = stack[0:top]
	}
	return res
}

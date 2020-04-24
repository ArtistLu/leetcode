/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func preorder(root *Node) []int {
	re := []int{}
	if root == nil {
		return re
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		re = append(re, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}
	return re
}
// @lc code=end


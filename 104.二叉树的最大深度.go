/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth(root *TreeNode) int {
	var re int
	if root == nil {
			return re
	}

	re++

	re += better(maxDepth(root.Left), maxDepth(root.Right))

	return re
}

func better(i,j int) int {
	if i> j {
			return i
	}

	return j
}
// @lc code=end


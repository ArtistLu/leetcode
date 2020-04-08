/*
 * @lc app=leetcode.cn id=1111 lang=golang
 *
 * [1111] 有效括号的嵌套深度
 */

// @lc code=start
func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))

	for i := range ans {
		if seq[i] == '(' {
			ans[i] = i & 1
		} else {
			ans[i] = (i+1) & 1
		}
	}

	return ans
}
// @lc code=end


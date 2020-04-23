/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	if len(S) == 0 {
		return S
	}
	re := ""
	var j int
	for i := 1; i < len(S); i++ {
		if S[i] == '(' {
			j++
		}
		if S[i] == ')' {
			j--
		}

		if j >= 0 {
			re += string(S[i])
		}

		if j == -1 {
			i++
			j = 0
		}
	}

	return re
}
// @lc code=end


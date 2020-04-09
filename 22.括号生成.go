/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	re := make([]string, 0)

	if n == 0 {
		return re
	}

	dfs("", n, n, &re)

	return re
}

func dfs(cur string, l, r int, re *[]string) {
	if l == 0 && r == 0 {
		*re = append(*re, cur)
	}

	if l > r {
		return
	}

	if l > 0 {
		dfs(cur+"(", l-1, r, re)
	}

	if r > 0 {
		dfs(cur+")", l, r-1, re)
	}
}
// @lc code=end


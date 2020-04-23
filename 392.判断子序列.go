/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	var j int
	for i := 0; i < len(t); i++ {
		if j == len(s)-1 && t[i] == s[j] {
			return true
		}
		if t[i] == s[j] {
			j++
		}
	}

	return false
}
// @lc code=end


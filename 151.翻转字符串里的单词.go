/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	sa := strings.Fields(s)
	for i, j := 0, len(sa)-1; i < j; i++ {
		sa[i], sa[j] = sa[j], sa[i]
		j--
	}
	return strings.Join(sa, " ")
}
// @lc code=end


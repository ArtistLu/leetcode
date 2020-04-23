/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sb := []byte(s)
	tb := []byte(t)
	bc := make([]int, 26)
	for i, c := range sb {
		bc[c-'a']++
		bc[tb[i]-'a']--
	}

	for _, c := range bc {
		if c != 0 {
			return false
		}
	}

	return true
}
// @lc code=end


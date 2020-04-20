/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")
	p := []byte(pattern)
	m := make(map[string]byte, len(p))
	u := make(map[byte]bool, 0)
	if len(s) != len(p) {
		return false
	}

	for i := 0; i < len(s); i++ {
		v, oks := m[s[i]]
		_, okp := u[p[i]]
		if oks && v != p[i] {
			return false
		}

		if !oks && okp {
			return false
		}

		m[s[i]] = p[i]
		u[p[i]] = true

	}
	return true
}
// @lc code=end


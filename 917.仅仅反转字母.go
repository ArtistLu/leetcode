/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(S string) string {
	s := []byte(S)
	i, j := 0, len(s)-1
	for i < j {
		for ; !isLetter(s[i]) && i < j; i++ {}
		for ; !isLetter(s[j]) && i < j; j-- {}
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return string(s)
}

func isLetter(l byte) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z')
}
// @lc code=end


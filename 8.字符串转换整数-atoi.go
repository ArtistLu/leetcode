/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	const (
		int32Max = 1<<31 - 1
		int32Min = -1 << 31
	)

	str = strings.TrimLeft(str, " ")

	op := 1
	re := 0
	for i, c := range str {
		if i == 0 && c == '-' {
			op = -1
		}

		if i == 0 && (c == '+' || c == '-') {
			continue
		}

		if c > '9' || c < '0' {
			break
		}

		re = re*10 + int(c-'0')

		if op*re > int32Max {
			return int32Max
		}

		if op*re < int32Min {
			return int32Min
		}

	}

	return op * re
}
// @lc code=end


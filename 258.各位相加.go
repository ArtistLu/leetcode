/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num < 10 {
		return num
	}

	re := 0
	for num > 0 {
		re += num % 10
		num /= 10
	}

	return addDigits(re)
}
// @lc code=end


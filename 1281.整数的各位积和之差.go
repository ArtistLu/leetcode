/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	x, a := 1, 0
	for n > 0 {
		u := n % 10
		x *= u
		a += u
		n = n / 10
	}
	
	return x - a
}
// @lc code=end


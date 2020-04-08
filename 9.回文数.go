/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	y := 0
	z := x
	for z >= 10 {
		y = y * 10 
		y += z % 10
		z /= 10
		 
	}

	return (y*10+z) == x
	
}
// @lc code=end


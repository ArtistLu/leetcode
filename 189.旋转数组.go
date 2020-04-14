/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	l := len(nums)

	var tmp int

	var cycles = gcd(l, k)
	var iters = l / cycles

	for startPos := 0; startPos < cycles; startPos++ {
		pos := startPos
		tmp = nums[pos]

		for i := 0; i < iters; i++ {
			next := (pos + k) % l

			nums[next], tmp = tmp, nums[next]

			pos += k
		}
	}
}

func gcd(a, b int) int {
	for {
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
}
// @lc code=end


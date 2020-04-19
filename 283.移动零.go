/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start

// 21/21 cases passed (4 ms)
// Your runtime beats 95.94 % of golang submissions
// Your memory usage beats 25 % of golang submissions (3.8 MB)
func moveZeroes(nums []int) {
	l := len(nums)
	for i, j := 0, 0; j < l; i, j = i+1, j+1 {
		for ; nums[j] == 0 && j < l-1; j++ {
		}
		nums[i] = nums[j]

		if i != j {
			nums[j] = 0
		}
	}
}

// 21/21 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 25 % of golang submissions (3.8 MB)
func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); i, j = i+1, j+1 {
		for ; nums[j] == 0 && j < len(nums)-1; j++ {
		}
		nums[i] = nums[j]

		if i != j {
			nums[j] = 0
		}
	}
}
// @lc code=end


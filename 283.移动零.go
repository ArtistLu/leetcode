/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
		l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] != 0 {
			continue
		}
		for j := i + 1; j < l; j++ {
			if nums[j] == 0 {
				continue
			}
			nums[j], nums[i] = nums[i], nums[j]
			break
		}
	}
}
// @lc code=end


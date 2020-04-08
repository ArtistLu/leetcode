/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	//1.找到数组中最大值的下标（假设为k）
	//2.遍历下标0-k，计算当前下标装水量
	//3.右边同左边

	// if len(height) == 0 {
	// 	return 0
	// }

	water := 0
	m := 0
	k := 0
	for i := range height {
		if height[i] > m {
			m = height[i]
			k = i
		}
	}

	cur := 0
	for i := 0; i < k; i++ {
		if cur > height[i] {
			water += (cur - height[i])
		} else {
			cur = height[i]
		}
	}

	cur = 0
	for i := len(height) - 1; i > k; i-- {
		if cur > height[i] {
			water += (cur - height[i])
		} else {
			cur = height[i]
		}
	}

	return water
}
// @lc code=end


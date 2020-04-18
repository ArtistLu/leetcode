/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index := target - nums[i]
		if b, ok := m[index]; ok {
			return []int{b, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
// @lc code=end


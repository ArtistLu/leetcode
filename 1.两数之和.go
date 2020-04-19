/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {

		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{}
}
// @lc code=end


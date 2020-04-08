/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, s, e int) {
	if s >= e {
		return
	}

	m := parting(nums, s, e)

	quickSort(nums, s, m-1)
	quickSort(nums, m+1, e)
}

func parting(nums []int, s, e int) int {
	tail := nums[e]

	i := s
	for j := s; j < e; j++ {
		if nums[j] < tail {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[e] = nums[e], nums[i]

	return i
}
// @lc code=end


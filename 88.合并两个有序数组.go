/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
		for j := m + i; j > 0; j-- {
			if nums1[j] < nums1[j-1] {
				nums1[j], nums1[j-1] = nums1[j-1], nums1[j]
			} else {
				break
			}
		}
	}
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	res := []int{}
	for _, v := range nums2 {
		if c, ok := m[v]; ok && c > 0 {
			m[v]--
			res = append(res, v)
		}
	}
	return res
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	length := len(A)
	re := make([]int, length)
	i, j, k := 0, length-1, length-1

	for i <= j {
		l := A[i] * A[i]
		r := A[j] * A[j]
		if l > r {
			re[k] = l
			i++
		} else {
			re[k] = r
			j--
		}
		k--
	}

	return re
}
// @lc code=end


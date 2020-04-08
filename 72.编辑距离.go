/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	if n == 0 || m == 0 {
		return n + m
	}

	matrix := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		matrix[0][i] = i
	}

	for j := 1; j <= m; j++ {
		matrix[j][0] = j
	}

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			last := matrix[j-1][i-1]
			if word1[i-1] != word2[j-1] {
				last++
			}
			matrix[j][i] = mdis(last, matrix[j-1][i]+1, matrix[j][i-1]+1)
		}
	}

	return matrix[m][n]
}

func mdis(i, j, k int) int {
	min := i
	if j < min {
		min = j
	}

	if k < min {
		min = k
	}

	return min
}
// @lc code=end


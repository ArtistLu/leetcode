/*
 * @lc app=leetcode.cn id=1162 lang=golang
 *
 * [1162] 地图分析
 *
 * https://leetcode-cn.com/problems/as-far-from-land-as-possible/description/
 *
 * algorithms
 * Medium (38.58%)
 * Likes:    40
 * Dislikes: 0
 * Total Accepted:    8.4K
 * Total Submissions: 20.9K
 * Testcase Example:  '[[1,0,1],[0,0,0],[1,0,1]]'
 *
 * 你现在手里有一份大小为 N x N 的『地图』（网格） grid，上面的每个『区域』（单元格）都用 0 和 1 标记好了。其中 0 代表海洋，1
 * 代表陆地，你知道距离陆地区域最远的海洋区域是是哪一个吗？请返回该海洋区域到离它最近的陆地区域的距离。
 * 
 * 我们这里说的距离是『曼哈顿距离』（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个区域之间的距离是 |x0 -
 * x1| + |y0 - y1| 。
 * 
 * 如果我们的地图上只有陆地或者海洋，请返回 -1。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：[[1,0,1],[0,0,0],[1,0,1]]
 * 输出：2
 * 解释： 
 * 海洋区域 (1, 1) 和所有陆地区域之间的距离都达到最大，最大距离为 2。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 输入：[[1,0,0],[0,0,0],[0,0,0]]
 * 输出：4
 * 解释： 
 * 海洋区域 (2, 2) 和所有陆地区域之间的距离都达到最大，最大距离为 4。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= grid.length == grid[0].length <= 100
 * grid[i][j] 不是 0 就是 1
 * 
 * 
 */

// @lc code=start
func maxDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	h, w := len(grid), len(grid[0])

	q := make([][2]int, 0)

	hasSea := false
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
			} else {
				hasSea = true
			}
		}
	}

	if hasSea == false || len(q) == 0 {
		return -1
	}

	arrived := make([][]int, h)

	for i := range arrived {
		arrived[i] = make([]int, w)
	}

	step := 0
	for len(q) > 0 {
		step++
		for range q {
			c := q[0]
			q = q[1:]
			for _, d := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				if i, j := c[0]+d[0], c[1]+d[1]; 0 <= i && i < h && 0 <= j && j < w {
					if grid[i][j] == 1 {
						continue
					}

					if arrived[i][j] == 0 {
						arrived[i][j] = step
						q = append(q, [2]int{i, j})
					}
				}
			}
		}
	}

	return step - 1
}
// @lc code=end


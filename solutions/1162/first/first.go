package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Print(maxDistance(grid))
}

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

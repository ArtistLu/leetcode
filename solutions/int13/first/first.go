package main

import "fmt"

func main() {
	test := [][3]int{{3, 1, 0}, {2, 3, 1}, {16, 8, 4}, {38, 15, 9}, {1, 1, 0}}
	for _, v := range test {
		re := movingCount(v[0], v[1], v[2])
		fmt.Print(re)
		fmt.Println()

	}
	// fmt.Print(re)
}

func movingCount(m int, n int, k int) int {
	re := 1
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	matrix[0][0] = 1

	q := [][2]int{{0, 0}}
	for len(q) > 0 {
		dot := q[0]
		q = q[1:]
		for _, step := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if i, j := step[0]+dot[0], step[1]+dot[1]; 0 <= i && i < m && 0 <= j && j < n {
				if matrix[i][j] != 1 && add(i)+add(j) <= k {
					re++
					q = append(q, [2]int{i, j})
				}
				matrix[i][j] = 1
			}
		}

	}

	return re
}

func add(s int) int {
	re := 0
	for s > 0 {
		re += s % 10
		s /= 10
	}
	return re
}

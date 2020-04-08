package main

import "fmt"

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// fmt.Print(board)

	// 	输入：
	// [
	//   [0,1,0],
	//   [0,0,1],
	//   [1,1,1],
	//   [0,0,0]
	// ]
	// 输出：
	// [
	//   [0,0,0],
	//   [1,0,1],
	//   [0,1,1],
	//   [0,1,0]
	// ]

	gameOfLife(board)
	fmt.Print(board)
}

func gameOfLife(board [][]int) {
	m := len(board)
	if m == 0 {
		return
	}

	n := len(board[0])
	if n == 0 {
		return
	}

	re := make([][]int, m)
	for i := range re {
		re[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for _, step := range [8][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}} {
				if r, c := i+step[0], j+step[1]; 0 <= r && r < m && 0 <= c && c < n {
					live += board[r][c]
				}
			}

			// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
			// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
			// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
			// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；

			re[i][j] = board[i][j]
			if live < 2 {
				re[i][j] = 0
			}

			if live == 3 && board[i][j] == 0 {
				re[i][j] = 1
			}

			if live > 3 {
				re[i][j] = 0
			}
		}
	}
	copy(board, re)
}

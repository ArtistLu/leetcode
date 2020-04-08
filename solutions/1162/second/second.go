package main

func main() {

}

func maxDistance(grid [][]int) int {
	max_dist := -1
	N := len(grid)
	for i, row := range grid {
		for j, g := range row {
			if g == 0 {
				dist := -1
				min_dist := -1
				for d_i := 0; d_i < N; d_i++ {
					for d_j := 0; d_j <= N; d_j++ {
						if d_i > 0 || d_j > 0 {
							sum := 0
							if i-d_i >= 0 {
								if j-d_j >= 0 {
									sum += grid[i-d_i][j-d_j]
								}
								if j+d_j < N {
									sum += grid[i-d_i][j+d_j]
								}
							}
							if i+d_i < N {
								if j-d_j >= 0 {
									sum += grid[i+d_i][j-d_j]
								}
								if j+d_j < N {
									sum += grid[i+d_i][j+d_j]
								}
							}

							if sum > 0 {
								dist = d_i + d_j
								break
							}
						}
					}
					if dist > 0 {
						if min_dist < 0 || dist < min_dist {
							min_dist = dist
						}
					}
				}
				if min_dist > max_dist {
					max_dist = min_dist
				}
			}
		}
	}
	return max_dist
}

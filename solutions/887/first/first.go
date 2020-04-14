package main

func main() {

}

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}

	for j := 1; j <= N; j++ {
		for i := 1; i <= K; i++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1] + 1
			if dp[i][j] >= N {
				return j
			}
		}
	}
	return N
}

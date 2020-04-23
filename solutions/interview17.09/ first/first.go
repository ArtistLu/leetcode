package main

import "fmt"

func main() {
	testMatrix := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for _, k := range testMatrix {
		fmt.Println(getKthMagicNumber(k))
	}
}

func getKthMagicNumber(k int) int {
	p3, p5, p7 := 0, 0, 0
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i < k; i++ {
		dp[i] = minNum(dp[p3]*3, minNum(dp[p5]*5, dp[p7]*7))

		if dp[p3]*3 == dp[i] {
			p3++
		}
		if dp[p5]*5 == dp[i] {
			p5++
		}
		if dp[p7]*7 == dp[i] {
			p7++
		}
	}
	return dp[len(dp)-1]
}

func minNum(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

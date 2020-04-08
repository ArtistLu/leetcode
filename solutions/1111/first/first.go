package main

import "fmt"

func main() {
	fmt.Print(2 & 1)
}

func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))

	for i := range ans {
		if seq[i] == '(' {
			ans[i] = i & 1
		} else {
			ans[i] = (i + 1) & 1
		}
	}

	return ans
}

package main

import "fmt"

func main() {
	test := []int{1, 3}
	for _, n := range test {
		re := generateParenthesis(n)
		fmt.Print(re)
	}
}

func generateParenthesis(n int) []string {
	re := make([]string, 0)

	if n == 0 {
		return re
	}

	dfs("", n, n, &re)

	return re
}

func dfs(cur string, l, r int, re *[]string) {
	if l == 0 && r == 0 {
		*re = append(*re, cur)
	}

	if l > r {
		return
	}

	if l > 0 {
		dfs(cur+"(", l-1, r, re)
	}

	if r > 0 {
		dfs(cur+")", l, r-1, re)
	}
}

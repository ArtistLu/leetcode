package main

import "fmt"

func main() {
	fmt.Print(lastRemaining(7, 7))
}

func lastRemaining(n int, m int) int {
	pos := 0
	for i := 2; i <= n; i++ {
		pos = (pos + m) % i
	}

	return pos
}

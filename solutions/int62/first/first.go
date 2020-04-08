package main

import "fmt"

func main() {
	fmt.Print(lastRemaining(10, 3))
}

//超时
func lastRemaining(n int, m int) int {
	q := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		q[i] = true
	}

	loop := 0
	step := 0
	end := 0
	for {
		i := loop % n
		if q[i] == true {
			step++
		}

		if q[i] == true && step%m == 0 {
			q[i] = false
			end++
		}

		if end == n-1 {
			break
		}

		loop++
	}

	for i, ok := range q {
		if ok == true {
			return i
		}
	}

	return 0
}

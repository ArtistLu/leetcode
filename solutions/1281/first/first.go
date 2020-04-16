package main

import "fmt"

func main() {
	testMatrix := []int{1, 11, 12, 124, 125}
	for _,v :=range testMatrix {
		fmt.Print(v, "->")
		re := subtractProductAndSum(v)
		fmt.Println(re)
	}
}

func subtractProductAndSum(n int) int {
	x, a := 1, 0
	for n > 0 {
		u := n % 10
		x *= u
		a += u
		n = n / 10
	}
	
	return x - a
}
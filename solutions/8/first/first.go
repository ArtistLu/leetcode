package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{
		"1",
		" -01",
		"- 9",
		"123",
	}

	for _, str := range strs {
		fmt.Print(myAtoi(str))
		fmt.Println()
	}
	// r := []byte(str)
	// str = []rune(str)
	// for _, n := range r {
	// 	// fmt.Print(n)
	// 	// fmt.Println()
	// 	// if n < 0 || n > 9 {
	// 	// 	break
	// 	// }
	// 	// re += re*10 + int(n)
	// }

	// fmt.Print(str[0], int(str[0]) > 9)

}

func myAtoi(str string) int {
	const (
		int32Max = 1<<31 - 1
		int32Min = -1 << 31
	)

	str = strings.TrimLeft(str, " ")

	op := 1
	re := 0
	for i, c := range str {
		if i == 0 && c == '-' {
			op = -1
		}

		if i == 0 && (c == '+' || c == '-') {
			continue
		}

		if c > '9' || c < '0' {
			break
		}

		re = re*10 + int(c-'0')

		if op*re > int32Max {
			return int32Max
		}

		if op*re < int32Min {
			return int32Min
		}

	}

	return op * re
}

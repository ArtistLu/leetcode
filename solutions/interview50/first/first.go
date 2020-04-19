package main

import "fmt"

func main() {
	testMatrix := []string{
		"aabcdef",
		"axyzzz",
		"",
	}

	for _, s := range testMatrix {
		first := firstUniqChar(s)
		fmt.Println(s, " -> ", string(first))
	}
}

func firstUniqChar(s string) byte {
	sb := []byte(s)
	mb := make(map[byte]int, len(sb))
	for _, c := range sb {
		mb[c]++
	}

	for _, c := range sb {
		if mb[c] == 1 {
			return c
		}
	}

	return ' '
}

package main

import "fmt"

func main() {
	testMatrix := [][]string{
		{"123", "023"},
		{"123", "123"},
		{"123", "024"},
	}

	for _, v := range testMatrix {
		fmt.Print(v[0], " | ", v[1], " -> ")
		fmt.Println(getHint(v[0], v[1]))
	}
}

func getHint(secret string, guess string) string {

	b := 0
	c := 0
	tmp := make([]int, 10)
	tmp[5]++
	tmp[5]--
	// fmt.Print(tmp)
	// return "ss"
	for i := 0; i < len(secret); i++ {
		s := secret[i] - '0'
		g := guess[i] - '0'
		if s == g {
			b++
			continue
		}
		if tmp[s] < 0 {
			c++
		}
		if tmp[g] > 0 {
			c++
		}
		tmp[s]++
		tmp[g]--
	}
	return fmt.Sprintf("%dA%dB", b, c)
}

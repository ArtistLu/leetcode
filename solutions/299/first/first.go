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
	s := []byte(secret)
	g := []byte(guess)
	m := make(map[byte]int, len(guess))
	gb := make([]byte, 0)
	a := 0
	b := 0
	for i := 0; i < len(s); i++ {
		if s[i] == g[i] {
			a++
		} else {
			gb = append(gb, g[i])
			m[s[i]]++
		}
	}

	for i := 0; i < len(gb); i++ {
		if k, ok := m[gb[i]]; k > 0 && ok {
			b++
			m[gb[i]]--
		}
	}

	return fmt.Sprintf("%dA%dB", a, b)
}

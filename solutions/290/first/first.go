package main

import (
	"fmt"
	"strings"
)

func main() {
	testMatrix := [][]string{
		{"abba", "dog cat cat dog"},
		{"abb", "dog cat cat dog"},
		{"abba", "dog cat cat fish"},
		{"abba", "dog dog dog dog"},
		{"abba", "dog cat cat fish"},
	}

	for _, v := range testMatrix {
		fmt.Print(v[0], "|", v[1], " -> ")
		fmt.Println(wordPattern(v[0], v[1]))
	}
}

func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")
	p := []byte(pattern)
	m := make(map[string]byte, len(p))
	u := make(map[byte]bool, 0)
	if len(s) != len(p) {
		return false
	}

	for i := 0; i < len(s); i++ {
		v, oks := m[s[i]]
		_, okp := u[p[i]]
		if oks && v != p[i] {
			return false
		}

		if !oks && okp {
			return false
		}

		m[s[i]] = p[i]
		u[p[i]] = true

	}
	return true
}

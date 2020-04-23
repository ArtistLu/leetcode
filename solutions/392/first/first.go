package main

import "fmt"

func main() {
	testMatrix := [][]string{
		{"abc", "aabdc"},
		{"", "aabdc"},
		{"acb", "ahbgdc"},
	}

	for _, arg := range testMatrix {
		fmt.Println(arg[0], "|", arg[1], "-> ", isSubsequence(arg[0], arg[1]))

	}
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	var j int
	for i := 0; i < len(t); i++ {
		if j == len(s)-1 && t[i] == s[j] {
			return true
		}
		if t[i] == s[j] {
			j++
		}
	}

	return false
}

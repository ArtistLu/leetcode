package main

import "fmt"

func main() {
	testMatrix := []string{
		"",
		"()",
		"()()(())",
		"(()())",
		"(()())(())",
	}

	for _, s := range testMatrix {
		fmt.Println(s, " -> ", removeOuterParentheses(s))
	}
}

func removeOuterParentheses(S string) string {
	if len(S) == 0 {
		return S
	}
	re := ""
	var j int
	for i := 1; i < len(S); i++ {
		if S[i] == '(' {
			j++
		}
		if S[i] == ')' {
			j--
		}

		if j >= 0 {
			re += string(S[i])
		}

		if j == -1 {
			i++
			j = 0
		}
	}

	return re
}

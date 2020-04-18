package main

import "fmt"

func main() {
	testMatrix := []string{
		"ab-cd",
		"aaabbb-",
		"728]",
	}

	for _, s := range testMatrix {
		fmt.Print(s, " -> ")
		fmt.Println(reverseOnlyLetters(s))
	}

}

func reverseOnlyLetters(S string) string {
	s := []byte(S)
	i, j := 0, len(s)-1
	for i < j {
		for ; !isLetter(s[i]) && i < j; i++ {
		}
		for ; !isLetter(s[j]) && i < j; j-- {
		}
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return string(s)
}

func isLetter(l byte) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z')
}

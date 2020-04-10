package main

import (
	"fmt"
	"strings"
)

func main() {
	re := reverseWords(" the good girl! ")
	fmt.Print(re)
}

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	sa := strings.Fields(s)
	for i, j := 0, len(sa)-1; i < j; i++ {
		sa[i], sa[j] = sa[j], sa[i]
		j--
	}
	return strings.Join(sa, " ")
}

package algorithms

import (
	"fmt"
	"strings"
)

func Repeated_string(s string, n int) int {
	fmt.Println("First ch: ",string(s[0:4]))
	sLen := len(s)
	if n < sLen {
		return strings.Count(s[0:n], "a")
	} else {
		//diffLen := n % sLen
	}
	return strings.Count(s, "a")
}

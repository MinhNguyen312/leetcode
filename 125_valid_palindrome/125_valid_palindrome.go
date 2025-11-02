package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// lowercase the string
	// strip string non-alphanumeric char from string
	//
	// This can be optimized in a different way
	// construct a reversed string
	// Traverse to compare

	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)
	r := regexp.MustCompile("[^a-zA-Z0-9]+")

	s = r.ReplaceAllString(s, "")

	// construct reversed string

	/*sb := strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}

	reversedString := sb.String()

	for i := 0; i < len(s); i++ {
		if s[i] != reversedString[i] {
			return false
		}
	}
	*/

	// use two pointer to compare
	// Faster and less memory
	i := 0
	j := len(s) - 1

	for i < len(s) && j >= 0 {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

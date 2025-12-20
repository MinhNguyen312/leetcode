package main

import "fmt"

// given two strings s, t,
// return true if t is and anagram of s
// false if otherwise

// Anagram definition:
// a word/phrase formed by rearranging the all letters of a different word/phrase. All are used exactly once
//
// Characteristics to check if anagram:
// - Length should be the same
// - letter used should be the same

func main() {
	s := "ratÖ"
	t := "tarÖ"

	fmt.Printf("t is anagram of s: %v\n", isAnagram(s, t))
}

func isAnagram(s string, t string) bool {

	// length check
	if len(t) != len(s) {
		return false
	}

	//	Create hashmap containing char and allow usage
	charMap := map[rune]int{}
	for _, v := range s {
		charMap[v]++
	}

	// char usage check
	for _, v := range t {
		charMap[v]--
		if charMap[v] < 0 {
			return false
		}
	}
	return true
}

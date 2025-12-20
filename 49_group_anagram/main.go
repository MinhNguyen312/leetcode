package main

import "fmt"

// Given array of strings strs, group anagrams together

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("%v\n", groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	result := [][]string{}

	strMap := make(map[[26]int][]string)
	// go through the array, check consecutives words
	for _, s := range strs {
		count := [26]int{}
		for _, c := range s {
			count[int(c)-int('a')] += 1
		}
		strMap[count] = append(strMap[count], s)
	}

	for _, v := range strMap {
		result = append(result, v)
	}

	return result
}

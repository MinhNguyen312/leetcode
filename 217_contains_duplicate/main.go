package main

import (
	"fmt"
	"sort"
)

// Given an int array 'nums' return 'true' if any value appears **at least twice** in the array
// return 'false' if every element is distinct

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Printf("Array %v contains duplicate: %v\n", arr, containsDuplicates(arr))

	arr = []int{1, 2, 3, 4}
	fmt.Printf("Array %v contains duplicate: %v\n", arr, containsDuplicates(arr))

	arr = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Printf("Array %v contains duplicate: %v\n", arr, containsDuplicates(arr))
}

// hashmap approach -> O(n)
// But requires more memory to create
//func containsDuplicates(nums []int) bool {
//	appearanceMap := map[int]int{}
//	for _, num := range nums {
//		if appearanceMap[num] == 0 {
//			appearanceMap[num] = 1
//		} else {
//			return true
//		}
//	}
//	return false
//}

// Faster and more memory optimized
func containsDuplicates(nums []int) bool {
	sort.Ints(nums)
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			return true
		}

		i++
	}
	return false
}

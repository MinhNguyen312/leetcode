package main

import "fmt"

func twosum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	// Return an empty slice if no solution is found
	return []int{}

}

func main() {
	nums1 := []int{2, 7, 11, 5}
	target1 := 9

	result := twosum(nums1, target1)

	fmt.Printf(`%d %d`, result[0], result[1])
}

package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2

Output: [1,2]

Example 2:

Input: nums = [1], k = 1

Output: [1]

Example 3:

Input: nums = [1,2,1,2,1,2,3,1,3,2], k = 2

Output: [1,2]

Constraints:

	1 <= nums.length <= 105
	-104 <= nums[i] <= 104
	k is in the range [1, the number of unique elements in the array].
	It is guaranteed that the answer is unique.
*/
func main() {
	nums := []int{5, -3, 9, 1, 7, 7, 9, 10, 2, 2, 10, 10, 3, -1, 3, 7, -9, -1, 3, 3}
	k := 3
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {

	// Get occurance of number

	// compare if occurance is greater then append to array, else keep on a different array.

	// greater will have an anchor point, will compare and do the same for the lesser number

	// join into fully sorted based on frequency

	// return from len(result) - k - 1 to len(result) -1

	freq := make(map[int]int)

	// populate map
	for _, v := range nums {
		freq[v]++
	}

	arr := make([][2]int, 0, len(freq))
	for num, cnt := range freq {
		arr = append(arr, [2]int{cnt, num})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	result := make([]int, k)
	for i := range k {
		result[i] = arr[i][1]
	}

	fmt.Println(result)

	return result
}

# Intuition
For first attempt, I tried to have an anchor string, then compare it to every other string in the array. If they are anagram, then add to a specific array.


This was not it



# Solution
The solution to this problem is actually a lot easier than the problem was gonna be. 

We can use an array to store the number of each character used. That will be the key of our hash-map. 

The map will look something like this
[0 1 1 0 0 0 0 1 0]: ['abc','bca']

What we are doing is using the number of character used for each string as the key, if the have the same key (i.e have the same character used to compose the string (anagram)), then will be added to the string array.

## Solution code
```Go
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
```
```
```

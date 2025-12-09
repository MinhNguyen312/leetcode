# Intuition
At first, I went straight for hashmap because:
- This would help me store the count of number
- Time complexity will be O(n)

But I simply forgot about space complexity
For this question, using a hashmap works but it would require computers to have memory overhead for adding numbers to the hashmap

# Solution
The best approach for this would be to sort it first, that way it will show if there are more than one identical number next to each other

```Go
	sort.Ints(nums)
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			return true
		}

		i++
	}
	return false
```


```

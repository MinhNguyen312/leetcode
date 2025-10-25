# Intuition
At first, I was going for comparing ASCII value (did not affect performance in anyway so I just use comparing character)

My intuition was to separate into valid and invalid cases
- Invalid cases
  - Length of string needs to be divisible by 2 (of course it's a pair of parentheses)
- Valid cases
  - If opening and closing parentheses is next to each other

I was going to use a open parentheses counter but this approach won't be able to handle things properly because it would only do the counting and not validating the parentheses.

In a similar fashion, we can use stack. 
For example:
  - Given this string `"({[]})"`, it should push all the opening parentheses into a stack (LIFO).
  - Then comparing the closing parentheses if they are a pair then will remove that character from the stack.
  - At the end, if there's no remaining open parentheses => return true

That was my intuition. However, there are a few cases that I missed.
- `"}("` : closing parentheses before having any opening. This causes panic because the program access the last element in the parentheses stack which only holds opening parentheses. (Index out of bound)
- `"([{}}])"`: Double closing. This also cause the same panic and error

# Solution
```
```go
func isValid(s string) bool {

	if len(s)%2 != 0 || s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}

	// check for valid
	// valid case
	var parentheses []rune

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			parentheses = append(parentheses, c)
			continue
		}

		switch c {
		case ')':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '(' {
				parentheses = parentheses[:len(parentheses)-1] // remove last element
			} else {
				return false
			}
		case ']':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '[' {
				parentheses = parentheses[:len(parentheses)-1] // remove last element
			} else {
				return false
			}
		case '}':
			if len(parentheses) > 0 && parentheses[len(parentheses)-1] == '{' {
				parentheses = parentheses[:len(parentheses)-1] // remove last element
			} else {
				return false
			}
		default:
			return false
		}
	}

	if len(parentheses) != 0 {
		return false
	}

	return true
}

```
```
```
```
```
# complexity
## Best case
O(n) - Because traverse string only one time

## Worst case
O(n) - All opening bracket is stored in stack

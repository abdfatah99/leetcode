# Longest Substring without repeating
[LeetCode Question](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

# Question

Given a string s, find the length of the longest substring without repeating characters.


Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

# Answer

> Open `Answer.excalidraw` for visual answer.

# Related Function in Go

## Slices

`slices` is a official package from go to work with array or slice. Based on the
`Answer.excalidraw` here are listed `slices` method that used to solve the problem.

### Contains

> Contains reports whether v is present in s.

**General Function**
```go
func Contains[S ~[]E, E comparable](s S, v E) bool
```

**Example Usage**
```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 1, 2, 3}
	fmt.Println(slices.Contains(numbers, 2)) // true
	fmt.Println(slices.Contains(numbers, 4)) // false
}
```

### Index

> Index returns the index of the first occurrence of v in s, or -1 if not present.

**General Function**
```go
func Index[S ~[]E, E comparable](s S, v E) int
```

**Example Usage**
```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, 8}
	fmt.Println(slices.Index(numbers, 8)) // 2
	fmt.Println(slices.Index(numbers, 7)) // -1 
}
```

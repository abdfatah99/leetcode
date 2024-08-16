package main

import (
	// "fmt"
	"math"

	slices "golang.org/x/exp/slices"
)

// using slice and slices.contains package to check the item
// O(n) for search
func LongestSubstring(param string) int {
	// pool := []string{}
	substring := []rune{}
	count := 0

	// create pool using hashmap
	// represent the collective of string with its amount of letter

	for _, letter := range param {
		// fmt.Printf("index: %d, letter: %c \n", index, letter)

		// [1] check letter occurence in substring
		//     if letter exist in the substring, add substring to pool, remove
		//     all item in the substring, then add the current letter as new
		//     element in the string
		//

		// if slices.Contains(substring, string(letter)) {
		// 	// [2] add substring to the pool
		// 	// [3] remove the substring
		// 	// [4] create the new substring based on current letter
		// 	// [5] continue the loop
		// 	fmt.Println("exist")
		// }

		index_occurence := slices.Index(substring, letter)

		// if the element is exist in the substring
		if index_occurence != -1 {

			substring = substring[index_occurence+1:]
		}

		substring = append(substring, letter)

		// increment number of existance of longest substring without repeating
		// example case.
		// if count 2, then new length without repeating is 3, then use 3 (max).
		// if count 3, then new length without repeaing is 1, then use 3 (max)
		count = int(math.Max(float64(count), float64(len(substring))))

	}

	// fmt.Println("result list: ", string(substring))
	return count
}

// using map to check the item in the list
// O(1)

package main

import "fmt"

// import "fmt"

func Twosum(nums []int, target int) []int {
	// var data int
	// for _, data := range nums {
	// 	fmt.Println(data)
	// }

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			// fmt.Println(nums[i] + nums[j])

			if nums[i]+nums[j] == target {
				result := []int{i, j}
				return result
			}
		}
	}

	return []int{0}
}

func TwoSumOptimze(nums []int, target int) []int {

	remainder := make(map[int]int)

	for index, data := range nums {
		// fmt.Println(index, data)
		// [loop 1]
		// - current (data): 2
		//   check if the current is number taht we search for
		//
		// ->   if value, exists := myMap[2]; exists {
		// ->       return myMap[7], index
		// ->       break;
		// ->   }
		//
		// - ad current remainder
		//   9 - 2 = 7 -> remainder = 7
		// - map -> remainder[remainder:index]
		// - map -> remainder[7:0]
		//
		// [loop 2]
		// - current (data): 7
		//   check if the current is number that we search for
		//
		// ->   if value, exists := myMap[7]; exists {
		// ->       return myMap[7], index
		// ->       break;
		// ->   }
		//
		// -  if not exist, add current remainder -> target - current
		//    9 - 7 = 2 -> remainder = 2
		// -  map -> remainder[7:0 2:1]

		// -> if data, exists := remainder[data] // we don't need the data
		if _, exists := remainder[data]; exists {
			return []int{remainder[data], index}
		}

		remainder[target-data] = index
	}

	return []int{0}
}

func main() {
	list1 := []int{2, 7, 11, 15}
	list2 := []int{3, 2, 4}
	list3 := []int{3, 3}
	fmt.Println(TwoSumOptimze(list1, 9))
	fmt.Println(TwoSumOptimze(list2, 6))
	fmt.Println(TwoSumOptimze(list3, 6))
}

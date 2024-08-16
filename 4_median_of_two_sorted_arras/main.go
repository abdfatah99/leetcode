package main

import "fmt"

// source: https://leetcode.com/problems/median-of-two-sorted-arrays/

func sort(nums []int) []int {
	// var result []int = []int{};

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
            if j == len(nums)-1 {
                continue
            }
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}

	return nums
}

/*
    Runtime: 43ms
    Memory: 5.35MB
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := append(nums1, nums2...)

	// sorting using has-table/hashmap/object?
    sorted := sort(result)

    // odd  = median: int -> (len(result)/2) + 1 -> 1 index
    // even = median: int -> (len(result)/2)     -> 2 index
    fmt.Printf("%d, type of result: %T \n", sorted, sorted)

    if len(sorted) % 2 == 0 {
        // this mean the length is even
        fIndex := (len(sorted)/2) - 1
        SIndex := fIndex + 1

        fmt.Println("first index and second index:")
        fmt.Println(sorted[fIndex], sorted[SIndex])

        return float64(float64(sorted[fIndex] + sorted[SIndex]) / 2.0)
    } else {
        fmt.Println("index of odd")
        fIndex := (len(sorted)/2) - 1

        fmt.Println(sorted[fIndex])

        return float64(sorted[fIndex+1])
    }
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("Median = %f \n\n", findMedianSortedArrays(nums1, nums2))

	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	fmt.Printf("Median = %f", findMedianSortedArrays(nums3, nums4))


}

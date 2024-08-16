package main

import "fmt"

func main() {
	list1 := []int{2, 4, 3}
	list2 := []int{5, 6, 4}
	linkedList1 := createLinkedList(list1)
	linkedList2 := createLinkedList(list2)

	result := addTwoNumbers(linkedList1.head, linkedList2.head)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

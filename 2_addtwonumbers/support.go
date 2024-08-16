package main

// create linked list from list of integer
func createLinkedList(list []int) *LinkedList {
	// first node is the first element of the list
	head := &Node{Val: list[0]}
	linkedList := &LinkedList{head: head}
	pointer := head

	// [1] skip the first element of the list
	// [2] add rest of the list as the tail
	for i := 1; i <= len(list)-1; i++ {
		pointer.Next = &Node{Val: list[i]}
		pointer = pointer.Next
	}

	return linkedList
}

func GetLinkedListValue(ll *LinkedList) []int {
	pointer := ll.head

	result := []int{}

	for pointer != nil {
		result = append(result, pointer.Val)
		pointer = pointer.Next
	}

	return result
}
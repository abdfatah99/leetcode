package main

import (
	"fmt"
	"reflect"
	"testing"
)



func TestCreateLinkedList(t *testing.T) {
	list1 := []int{2, 4, 3}
	// list2 := []int{5, 6, 4}

	linkedList := createLinkedList(list1)
	llValue := GetLinkedListValue(linkedList)
	fmt.Println("linked list value:", llValue)

	if reflect.DeepEqual(llValue, list1) {
		t.Logf("test pass")
	} else {
		t.Error("Linked List value is not the same")
	}
}

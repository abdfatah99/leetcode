package main

func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	// 1. Traverse the linked list
	// 2. Add each node value while traverse it

	result := &Node{}

	pointer := result
	//		first pointer point to the first initialize node. This first init node
	// 		will become the head of the linked list.
	// 		looping over the longest linked list and the calculation result becomes
	// 		the tail of the head and soon.

	carry := 0
	// 		carry of the calculation

	// loop over the longest linked list
	for l1 != nil || l2 != nil {
		
		// initialize value for calculation each node
		var l1Val, l2Val int

		// check for each list-node if the list node is not nil, then get the value
		// and calculate the value
		if l1 != nil {
			l1Val = l1.Val

			// then set the node to the next node
			l1 = l1.Next
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		// calculate each
		sum := l1Val + l2Val + carry
		carry = sum / 10

		// create the tail of linked-list
		pointer.Next = &Node{Val: sum % 10}

		// move the pointer to the tail
		pointer = pointer.Next
	}

	// this is the last stage if after all calculation exist a carry, create
	// the carry as the last tail.
	if carry > 0 {
		pointer.Next = &Node{Val: carry}
	}

	// pass the result after the init node, which the head
	return result.Next
}

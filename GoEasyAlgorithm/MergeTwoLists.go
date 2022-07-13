package main

// Declare a struct
type ListNode struct{
	// Data
	Value int
	// Pointer
	Next *ListNode
}

func mergeTwoLists(L1 *ListNode, L2 *ListNode) *ListNode {
	// Declare a head pointer and declare an answer list
	head := &ListNode{}
	currentNode := head

	// For to judge every node
	for L1 != nil || L2 != nil {
		// Compares the size of two non-empty nodes
		if L1 != nil && L2 != nil {
			if L1.Value < L2.Value {
				// Modify current node next pointer to point to more less than other
				currentNode.Next = L1
				L1 = L1.Next
			}else {
				currentNode.Next = L2
				L2 = L2.Next
			}
			currentNode = currentNode.Next
		}else if L1 != nil {
			currentNode.Next = L1
			break
		}else {
			currentNode.Next = L2
			break
		}
	}
	return head.Next
}
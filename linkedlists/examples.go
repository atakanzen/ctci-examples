package linkedlists

import (
	ll "github.com/atakanzen/go-dsa/linkedlist"
)

// 2.1 Write code to remove duplicates from an unsorted linked list. Follow up, how about without an intermediary buffer

func RemoveDups(linkedList *ll.DoublyLinkedList) []interface{} {
	frequencyMap := make(map[interface{}]int, linkedList.Length) // Additional intermediary buffer

	n := linkedList.Head

	for i := 0; n != nil; i++ {

		if _, ok := frequencyMap[n.Value]; ok {
			linkedList.RemoveAt(i)
			i -= 1 // Since ith has been removed we decrement to align with new length
		} else {
			frequencyMap[n.Value] = 1
		}
		n = n.Next
	}

	return linkedList.Values()
}

func RemoveDupsWithoutBuffer(linkedList *ll.DoublyLinkedList) []interface{} {

	n := linkedList.Head
	for i := 0; i < linkedList.Length; i++ {
		m := n.Next
		for j := i + 1; j < linkedList.Length; j++ {
			// Check
			if n.Value == m.Value {
				linkedList.RemoveAt(j)
			}
			m = m.Next
		}
		n = n.Next
	}

	return linkedList.Values()
}

// 2.2 Implement an algorithm to find kth to last element of a singly linked list

func KthToLastElement(head *ll.Node, k int) interface{} {
	length := countLinkedList(head) // O(N)

	if k > length-1 || k <= 0 {
		return nil
	}

	kTh := length - k

	n := head
	for i := 0; n != nil; i++ {
		if i == kTh {
			return n.Value
		}
		n = n.Next
	}

	return nil
}

func countLinkedList(head *ll.Node) int {
	n := head
	count := 0
	for n != nil {
		count++
		n = n.Next
	}

	return count
}

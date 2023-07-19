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

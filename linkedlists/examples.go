package linkedlists

import (
	"fmt"

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

func KthToLastElementRecursive(head *ll.Node, k int, i *int) interface{} {
	if head == nil {
		return nil
	}

	val := KthToLastElementRecursive(head.Next, k, i)
	*i += 1

	if *i == k {
		return head.Value
	}

	return val
}

func KthToLastElementRunner(head *ll.Node, k int) interface{} {
	if k < 0 {
		return nil
	}

	p1 := head
	p2 := head

	// Moving p1 k items forward
	for i := 0; i < k; i++ {
		if p1 == nil { // out of bounds
			return nil
		}

		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2.Value
}

// 2.3 Delete Middle Node: Implement an algorithm to delete a node in the middle (any value other than the first and last node) of a singly linked list given only access to that node

func DeleteMiddleNode(node *ll.Node) bool {
	if node == nil || node.Next == nil {
		return false // Tail Node
	}

	node.Next = node.Next.Next
	return true
}

// 2.4
// TODO
func Partition(node *ll.Node, pivot interface{}) *ll.Node {
	head := node
	tail := node

	for node != nil {
		fmt.Println(node.Value)
		next := node.Next
		if node.Value.(int) < pivot.(int) {
			head.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}
		node = next
	}

	return head
}

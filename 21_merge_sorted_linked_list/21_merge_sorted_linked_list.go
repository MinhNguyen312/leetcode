package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	// Plan
	// nil check
	// list1[0].Next == nil {}
	// if list1[0] > list2[0] {result = list2[0]; list2[0].Next = list1[0]}

	dummy := ListNode{}
	tail := &dummy

	for list1 != nil && list2 != nil {
		fmt.Printf("list1: %d\tlist2: %d\n", list1.Val, list2.Val)

		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next

	}

	if list1 == nil {
		tail.Next = list2
	}

	if list2 == nil {
		tail.Next = list1
	}

	return dummy.Next
}

func main() {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 4,
	}

	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{
		Val: 1,
	}
	node5 := &ListNode{
		Val: 3,
	}
	node6 := &ListNode{
		Val: 4,
	}
	node7 := &ListNode{
		Val: 5,
	}
	node8 := &ListNode{
		Val: 6,
	}

	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8

	result := mergeTwoList(node1, node4)
	current := result
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}

}

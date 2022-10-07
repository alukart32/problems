// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
// Input: list1 = [], list2 = []
// Output: []
//
// Input: list1 = [], list2 = [0]
// Output: [0]
package main

import (
	"fmt"
)

func main() {

	//	[5]
	//	[1,2,4]

	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l := mergeTwoLists(list1, list2)
	fmt.Println(l)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// set head
	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	var p *ListNode
	p = head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}

	return head
}

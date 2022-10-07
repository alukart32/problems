// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Examples:
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
//
// Input: l1 = [0], l2 = [0]
// Output: [0]
//
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	shift := 0
	l3 := new(ListNode)
	cur := l3
	for l1 != nil || l2 != nil || shift != 0 {
		val1, val2 := 0, 0
		if l1 != nil {
			val1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			val2, l2 = l2.Val, l2.Next
		}
		sum := val1 + val2 + shift
		// get last digit of xx number
		shift = sum / 10
		cur.Next = &ListNode{
			Val: sum % 10, // set number < 10 or xx % 10
		}
		cur = cur.Next
	}
	return l3.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	hl3 := l3
	for l3 != nil {
		if l1 != nil {
			l3.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l3.Val += l2.Val
			l2 = l2.Next
		}
		if l3.Val >= 10 {
			l3.Val -= 10
			l3.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			l3.Next = &ListNode{Val: 0}
		}
		l3 = l3.Next
	}
	return hl3
}

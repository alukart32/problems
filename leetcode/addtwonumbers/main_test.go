package main

import (
	"errors"
	"testing"
)

func TestAddTwoNumbers_342_465(t *testing.T) {
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	wanted := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}

	actual := addTwoNumbers(list1, list2)
	if err := isListsEqual(actual, wanted); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers_0_0(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{}
	wanted := &ListNode{}

	actual := addTwoNumbers(list1, list2)
	if err := isListsEqual(actual, wanted); err != nil {
		t.Error(err)
	}
}

func TestAddTwoNumbers_9999999_9999(t *testing.T) {
	// 9,9,9,9,9,9,9
	list1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}
	list2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	// 8,9,9,9,0,0,0,1
	wanted := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 1},
							},
						},
					},
				},
			},
		},
	}

	actual := addTwoNumbers(list1, list2)
	if err := isListsEqual(actual, wanted); err != nil {
		t.Error(err)
	}
}

func isListsEqual(pa, pw *ListNode) error {
	for pa != nil && pw != nil {
		if pa.Val != pw.Val {
			return errors.New("wanted and actual numbers are not equal")
		}
		pa = pa.Next
		pw = pw.Next
	}
	return nil
}

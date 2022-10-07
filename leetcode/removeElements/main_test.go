package main

import (
	"testing"
)

func TestRemoveElements1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expectedNums := []int{2, 2}

	k := removeElement(nums, val)

	if k != len(expectedNums) {
		t.Fatalf("different length: %d != %d", len(expectedNums), k)
	}

	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			t.Fatal("arrays not equal")
		}
	}
}

func TestRemoveElements2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 4, 2, 6, 9, 2}
	val := 2
	expectedNums := []int{0, 1, 4, 6, 9}

	k := removeElement(nums, val)

	if k != len(expectedNums) {
		t.Fatalf("different length: %d != %d", len(expectedNums), k)
	}

	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			t.Fatal("arrays not equal")
		}
	}
}

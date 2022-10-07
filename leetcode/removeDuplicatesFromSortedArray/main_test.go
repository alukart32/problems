package main

import "testing"

func TestRemoveDUplicatesFromArray1(t *testing.T) {
	nums := []int{0, 0, 1, 2, 2, 2, 3, 4, 5, 5, 6}
	expectedNums := []int{0, 1, 2, 3, 4, 5, 6}

	k := removeDuplicates(nums)

	if k != len(expectedNums) {
		t.Fatalf("different length between k and expectedNums: %d != %d", k, len(expectedNums))
	}
	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			t.Fatal("arrays not equal")
		}
	}
}

func TestRemoveDUplicatesFromArray2(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0}
	expectedNums := []int{0}

	k := removeDuplicates(nums)

	if k != len(expectedNums) {
		t.Fatalf("different length between k and expectedNums: %d != %d", k, len(expectedNums))
	}
	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			t.Fatal("arrays not equal")
		}
	}
}

func TestRemoveDUplicatesFromArray3(t *testing.T) {
	nums := []int{0, 0, 0, 0, 0, 1, 2, 2}
	expectedNums := []int{0, 1, 2}

	k := removeDuplicates(nums)

	if k != len(expectedNums) {
		t.Fatalf("different length between k and expectedNums: %d != %d", k, len(expectedNums))
	}
	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			t.Fatal("arrays not equal")
		}
	}
}

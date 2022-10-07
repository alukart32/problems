package main

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	x := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[x] = nums[i]
			x++
		}
	}
	return x
}

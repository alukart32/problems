package main

func removeElement(nums []int, val int) int {
	if nums == nil {
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[x] = nums[i]
			x++
		}
	}
	return x
}

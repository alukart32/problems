// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
//
// Example:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

// Solution
//
//	left_idx		  right_idx  target
//
// [ 2,    7,	11,   15]  	     9
//
// if (left_indx + right_indx > target)
//
//	right_index--
//
// else (left_indx + right_indx < target)
//
//	left_index++
//
// else (left_indx + right_indx = target)
//
//	indices: left_idx & right_idx
//
// go until left < right
func twoSumWithOrderinfElem(nums []int, target int) []int {
	// prepare slice for work
	sort.Ints(nums)
	// algorithm to find a pair of indices of the two numbers such that they add up to target
	leftIdx := 0
	rightIdx := len(nums) - 1
	for leftIdx < rightIdx {
		if nums[leftIdx]+nums[rightIdx] > target {
			rightIdx--
		} else if nums[leftIdx]+nums[rightIdx] < target {
			leftIdx++
		} else if nums[leftIdx]+nums[rightIdx] == target {
			return []int{leftIdx, rightIdx}
		}
	}
	return []int{}
}

// [2,7,11,15], target = 9
// 9 - 2 = 7
// 9 - 7 = 2
// 9 - 11 = -2
// 9 - 15 = -6
//
// # Solution with map
//
// 1) map[int]int, where key is i and value is target - nums[i]
// 2) before adding a new key-value pair searching possible key like target - nums[i] in map
func twoSum(nums []int, target int) []int {
	sumMap := map[int]int{}
	for i := range nums {
		if _, ok := sumMap[nums[i]]; ok {
			return []int{sumMap[nums[i]], i}
		}
		sumMap[target-nums[i]] = i
	}
	return []int{}
}

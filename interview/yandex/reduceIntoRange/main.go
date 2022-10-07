// Дан список интов, повторяющихся элементов в списке нет.
// Нужно преобразовать это множество в строку, сворачивая соседние по числовому ряду числа в диапазоны.
//
// Примеры:
// [1, 4, 5, 2, 3, 9, 8, 11, 0] => "0-5, 8-9, 11"
// [1, 4, 3, 2] => "1-4"
// [1, 4] => "1, 4"
package main

import (
	"sort"
	"strconv"
	"strings"
)

func reuceInroRange(nums ...int) string {
	if len(nums) == 0 {
		return ""
	}

	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	var sb strings.Builder
	start, end := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if start == end {
				sb.WriteString(strconv.Itoa(nums[start]) + ", ")
				start++
			} else {
				sb.WriteString(strconv.Itoa(nums[start]) + "-" + strconv.Itoa(nums[end]) + ", ")
				start = end + 1
			}
		}
		end++
	}

	// check last number case
	// is it in a range or single
	if start == end {
		sb.WriteString(strconv.Itoa(nums[start]))
	} else {
		sb.WriteString(strconv.Itoa(nums[start]) + "-" + strconv.Itoa(nums[end]))
	}

	return sb.String()
}

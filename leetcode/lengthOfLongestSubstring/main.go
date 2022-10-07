// Given a string s, find the length of the longest substring without repeating characters.
//
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
package main

//  slide window
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength, start, end := 0, 0, 0
	uniqueCharacters := make(map[rune]bool)
	for end < len(s) {
		if _, ok := uniqueCharacters[rune(s[end])]; ok {
			delete(uniqueCharacters, rune(s[start]))
			start++
		} else {
			uniqueCharacters[rune(s[end])] = true
			end++
			if len(uniqueCharacters) > maxLength {
				maxLength = len(uniqueCharacters)
			}
		}
	}
	return maxLength
}

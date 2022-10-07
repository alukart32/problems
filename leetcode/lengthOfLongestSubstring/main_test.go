package main

import "testing"

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		value string
		want  int
	}{
		{value: "abcabcbb", want: 3},
		{value: "bbbbbb", want: 1},
		{value: "pwwkew", want: 3},
	}

	for _, tc := range testCases {
		t.Run(tc.value, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tc.value)
			if actual != tc.want {
				t.Errorf("longest substring for %s has length %d, but get length %d", tc.value, tc.want, actual)
			}
		})
	}
}

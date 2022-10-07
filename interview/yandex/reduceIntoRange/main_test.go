package main

import "testing"

func TestReduceIntoRange(t *testing.T) {
	testCases := []struct {
		input []int
		want  string
	}{
		{
			input: []int{1, 4, 5, 2, 3, 9, 8, 11, 0},
			want:  "0-5, 8-9, 11",
		},
		{
			input: []int{1, 4, 3, 2},
			want:  "1-4",
		},
		{
			input: []int{1, 4},
			want:  "1, 4",
		},
		{
			input: []int{1, 2, 4, 6, 7, 9},
			want:  "1-2, 4, 6-7, 9",
		},
		{
			input: []int{1, 3, 5, 6, 8, 10},
			want:  "1, 3, 5-6, 8, 10",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			actual := reuceInroRange(tc.input...)
			if actual != tc.want {
				t.Errorf("for %v expected %s, get %v", tc.input, tc.want, actual)
			}
		})
	}
}

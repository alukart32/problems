package main

import "testing"

func TestLastWordLen(t *testing.T) {
	s := "my suprer text  "
	l := 4

	if lengthOfLastWord(s) != l {
		t.Fail()
	}
}

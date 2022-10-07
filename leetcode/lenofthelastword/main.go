package main

import "unicode"

func lengthOfLastWord(s string) int {
	x := 0

	// some__text___
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && x != 0 {
			break
		}
		if s[i] != ' ' {
			x++
		}
	}
	return x
}

func lengthOfLastWord2(s string) int {
	pos := len(s) - 1
	for ; pos >= 0; pos-- {
		if !unicode.IsSpace(rune(s[pos])) {
			break
		}
	}
	s = s[:pos+1]
	if len(s) == 0 {
		return 0
	}

	for ; pos >= 0; pos-- {
		if unicode.IsSpace(rune(s[pos])) {
			break
		}
	}
	return len(s[pos+1:])
}

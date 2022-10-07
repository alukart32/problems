/*
rules:
- the same character means '+' operation: VVV = V + V + V = 5 + 5 + 5 = 15
- I before V or X means '-' operation to them: IV : = V - I = 5 - 1 = 4
- X before L or C means '-' operation to them
- C before D or M means '-' operation to them

solution:
1) give all charaters some weight: I = 1, V = 5, X = 10, L = 50, C = 100, D = 500, M = 1000
2) if cur_char < next_char then base += next_char - cur_char: MCMXCIV = 1000 + (1000 - 100) + (100 - 10) + (5 - 1) = 1994
*/
package main

import "strings"

func main() {
	//println(romanToInt("MCDLXXVI"))
	println(intToRoman(8567))
}

var romanNumeralMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if romanNumeralMap[s[i]] < romanNumeralMap[s[i+1]] {
				sum -= romanNumeralMap[s[i]]
			} else {
				sum += romanNumeralMap[s[i]]
			}
		} else {
			sum += romanNumeralMap[s[i]]
		}
	}

	return sum
}

func intToRoman(num int) string {
	romanNumeralMapping := []struct {
		value   int
		numeral string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	convertedNumeral := ""
	for _, conversion := range romanNumeralMapping {
		symbCount := num / conversion.value
		if symbCount >= 1 {
			num %= conversion.value
			convertedNumeral += strings.Repeat(conversion.numeral, symbCount)
		}
	}
	return convertedNumeral
}

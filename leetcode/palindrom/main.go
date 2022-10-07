package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", isPalindrome(9889))
}

func isPalindrome(x int) bool {
	return recursive(x, &x)
}

// dup - указатель на оригинальное значение
// v - лоакальная копия переданного значения
func recursive(v int, dup *int) bool {
	if v < 0 {
		return false
	}
	if v < 10 {
		return true
	}

	// пытаемся дойти до с конца до самого первого числа
	// , чтобы проверять от него первую и вторую часть числа
	//     >       >
	//   a b c | c f g
	flag := recursive(v/10, dup)

	*dup = *dup / 10
	lastElem := v % 10

	if flag && lastElem == *dup%10 {
		return true
	}
	return false
}

func isPalindrom2(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func isPlaindrom3(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	p := 0
	for x > 0 {
		tmp := x % 10
		x = x / 10
		p *= 10
		p += tmp
	}
	return origin-p == 0
}

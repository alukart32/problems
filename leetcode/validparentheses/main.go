/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
package main

import "fmt"

func main() {
	println(isValid2("(([]))"))
}

// Решение проще
// Пушим в slice-stack только закрывающие скобки, и если встречается не открывающая
// , то смотрим на валидность закрывающей скобки
func isValid2(s string) bool {
	stack := make([]byte, 0, (len(s)/2)+1)
	for i := range s {
		switch {
		case s[i] == '(':
			stack = append(stack, ')')
		case s[i] == '{':
			stack = append(stack, '}')
		case s[i] == '[':
			stack = append(stack, ']')
		case len(stack) == 0 || s[i] != stack[len(stack)-1]:
			return false
		default:
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := Stack{}
	// без доп. копирования символов
	for i := range s {
		switch s[i] {
		case '(':
			stack.Push(s[i])
		case '[':
			stack.Push(s[i])
		case '{':
			stack.Push(s[i])
		case ')':
			el, _ := stack.Top()
			if el == '(' {
				stack.Pop()
			} else {
				return false
			}
		case '}':
			el, _ := stack.Top()
			if el == '{' {
				stack.Pop()
			} else {
				return false
			}
		case ']':
			el, _ := stack.Top()
			if el == '[' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

// thread-unsafe
type Stack struct {
	items []byte
}

func (s *Stack) Push(el byte) {
	s.items = append(s.items, el)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (byte, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

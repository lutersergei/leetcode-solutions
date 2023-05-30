package main

import "fmt"

// Valid Parentheses
func main() {
	fmt.Println(isValid("{}()[]"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	var stack []byte
	a := byte('[')
	A := byte(']')
	b := byte('{')
	B := byte('}')
	c := byte('(')
	C := byte(')')

	m := make(map[byte]byte)
	m[A] = a
	m[B] = b
	m[C] = c

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case a, b, c:
			stack = append(stack, s[i])
		case A, B, C:
			if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

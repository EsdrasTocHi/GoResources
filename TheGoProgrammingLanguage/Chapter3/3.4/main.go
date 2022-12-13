package main

import "fmt"

// if the answer is already determined by the value of the left operand, the right operand is not evaluated

func main() {
	s := ""

	// this not produce and error because the second
	// operand is not evaluated
	if s != "" && s[0] == 'x' {
		fmt.Println("hello world")
	}
}

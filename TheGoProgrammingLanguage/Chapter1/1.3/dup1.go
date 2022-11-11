package main

import (
	"bufio"
	"fmt"
	"os"
)

// What is a map?
// Do you know the JSON format? Well... similar data structure with it.
// The maps have a key and value, you can access to the value using the key.
// When you create a map you must define the type of the key and the type of the value and
// you can use as a key any data type that you can compare with the == operand.

// For example:
// If you declare a map with... map := make(map[int]bool)
// the type between the square brackets is the key value (string) and the type next to the key type
// is the value type
// when you add a new elements, your map can look like this:
// { 10 : true, 5 : false }
// this map has two values.
// if you need a value of the key 10, you can acces to it with map[10], and this will return true

// if you want to iterate a map you can use range, this will return the key and the value respectively

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			/*
				%d decimal integer
				%x, %o, %b integer in hexadecimal, octal, binary
				%f, %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00
				%t boolean: true or false
				%c rune (Unicode code point)
				%s string
				%q quoted string "abc" or rune 'c'
				%v any value in a natural format
				%T type of any value
				%% literal percent sign (no operand)
			*/
		}
	}
}

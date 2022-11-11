package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var finalText, separator string
	//other declaration forms:
	/*

		var finalText = ""
		finalText := ""
		var finalText string = ""

	*/

	//os.Args is a string slice, this contains the params that you put in the command line
	//when you run your program

	//Example:
	//if you run your code with:
	// 	go run main.go param1 param2 ... paramN
	//os.Args must be ["param1" , "param2" , "..." , "paramN"]
	for i := 1; i < len(os.Args); i++ {
		finalText += separator + os.Args[i]
		separator = " "
	}
	fmt.Println(finalText)

	separator = ""
	finalText = ""
	//other form to iterate over a range of string or slice

	//range returns two values: index and the value of the element.
	// in this case we use _ to ignore the index value because we won't use it
	for _, arg := range os.Args[1:] {
		finalText += separator + arg
		separator = " "
	}

	fmt.Println(finalText)

	// the most efficient form to do an echo function:
	fmt.Println(strings.Join(os.Args[1:], " "))
}

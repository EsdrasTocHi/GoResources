package main

import "fmt"

// you can use the math/complex package for working with complex numbers

func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = 3 + 4i        // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"
}

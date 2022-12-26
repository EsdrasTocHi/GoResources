package main

import "fmt"

func main() {
	var arr1 [3]int
	var arr2 = [...]int{1, 2, 3, 4, 5}
	// [5]int type != [3]int type

	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)

	r := [...]int{9: 0} // the first value is the size of the array -1, the secons is the last value
	fmt.Println(r)
}

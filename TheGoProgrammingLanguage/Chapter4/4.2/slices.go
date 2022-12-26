package main

import "fmt"

// slices has three components: pointer, length and capacity
// POINTER:
//		points to the first element of the array that is reachable through the slice
//		which is not necesarily the first element of the array
// LENGTH:
//		number of slice elements. This can not be greater than the capacity
// CAPACITY:
//		is usually the number of elements between the start of the slice and the end
//		of the underlying array

func main() {
	months := [...]string{1: "January", 2: "february", 3: "march", 4: "april", 5: "may", 6: "june", 7: "july", 8: "august", 9: "september", 10: "october", 11: "november", 12: "December"}
	fmt.Println(months)
	fmt.Printf("length: %d   capacity: %d\n", len(months), cap(months))

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2) // ["April" "May" "June"]
	fmt.Printf("length: %d   capacity: %d\n", len(Q2), cap(Q2))
	fmt.Println(summer) // ["June" "July" "August"]
	fmt.Printf("length: %d   capacity: %d\n", len(summer), cap(summer))

	// Slicing beyond cap(s) causes a panic, but slicing beyond len(s) extends the slice, so the result may be longer than the original:

	//fmt.Println(summer[:20])    // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
	fmt.Printf("length: %d   capacity: %d\n", len(endlessSummer), cap(endlessSummer))

	slice1, slice2 := []int{1, 2, 3}, []int{4, 5}

	// when the base array dont have enough capacity, this duplicates it
	slice3 := append(slice1, slice2...)

	fmt.Println(slice3)
	fmt.Printf("length: %d   capacity: %d\n", len(slice3), cap(slice3))

}

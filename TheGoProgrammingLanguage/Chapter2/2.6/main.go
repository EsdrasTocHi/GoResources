package main

import (
	"Temp/popcount"
	"Temp/tempconv"
	"fmt"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC).String())
	fmt.Print("Popcount ->   ")
	fmt.Println(popcount.PopCount(100))
}

package main

import "fmt"

// in this case, float64 determines the structure and representation
// of Celsius and Farenheit
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// We can create methods to our types
func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func (f Fahrenheit) ToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := BoilingC.ToF()
	fmt.Printf("%g\n", boilingF-FreezingC.ToF()) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	//fmt.Println(c == f) // missmatch
	fmt.Println(c == Celsius(f)) //true
}

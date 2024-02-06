package main

import (
	"fmt"
	"math"
)

//Arithemetic operations
func arithemetic_example() {
	price, tax := 275.00, 27.40
	sum := price + tax
	differnce := price - tax
	product := price * tax
	quotient := price / tax

	var intVal = math.MaxInt64
	var floatVal = math.MaxFloat64

	fmt.Println(sum)
	fmt.Println(differnce)
	fmt.Println(product)
	fmt.Println(quotient)
	fmt.Println(intVal * 2)
	fmt.Println(floatVal * 2)
	fmt.Println(math.IsInf((floatVal * 2), 0))
}

//Perfoming Explicit Type Conversions
func type_conversion_example() {
	price := 240
	tax := 20.00
	fmt.Println(float64(price) + tax)
}

func main() {
	type_conversion_example()
}

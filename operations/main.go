package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, Operations")
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

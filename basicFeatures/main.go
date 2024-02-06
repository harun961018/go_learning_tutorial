package main

import (
	"fmt"
	"math/rand"
)

func const_example() {
	const price float32 = 275.00
	const tax float32 = 27.50
	fmt.Println(price + tax)
}

func rand_library_example() {
	fmt.Println(rand.Int())
	fmt.Println("value:", rand.Int())
}

func main() {
	const_example()
}

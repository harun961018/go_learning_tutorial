package main

import (
	"fmt"
	"math/rand"
	"sort"
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

//Using an Untyped Constant
func untyped_const_example() {
	const price float32 = 275.00
	const tax float32 = 25.00
	const quantity = 2
	fmt.Println("Total", quantity*(price+tax))
}

//Defining Multiple Constants with a Single Statement
func define_multiple_constant() {
	const price, tax float32 = 275.00, 25.00
	const quantity, inStock = 2, true
	fmt.Println("Total:", quantity*(price+tax))
	fmt.Println("In stock", inStock)
}

//Revisiting Literal Values
func literal_value() {
	const price, tax float32 = 275.00, 25.00
	const quantity, instock = 2, true
	fmt.Println("Total:", 2*quantity*(price+tax))
	fmt.Println("In stock", instock)
}

// Definint variable
func define_variable() {
	var price float32 = 275.00
	var tax float32 = 27.00
	fmt.Println("total-1:", price+tax)
	price = 300
	fmt.Println("total-2:", price+tax)
}

//Redefine variable using short syntax
func redefine_variable() {
	price, tax := 10.00, 10.00
	fmt.Println(price + tax)
	price2, tax := 30.00, 20.00
	fmt.Println(price2 + tax)
}

//Using the Blank Identifier
func blank_identifier() {
	price, tax, inStock, _ := 275.00, 25.00, true, true
	var _ = "Alice"
	fmt.Println("Total:", price+tax)
	fmt.Println("In stock", inStock)
}

//Using pointer
func pointer_example() {
	names := [3]string{"Alice", "bob", "Chile"}
	secondPosition := &names[1]
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(*secondPosition)
}

func main() {
	// define_multiple_constant()
	// literal_value()
	// define_variable()
	// redefine_variable()
	// blank_identifier()
	pointer_example()
}

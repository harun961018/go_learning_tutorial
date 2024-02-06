package main

import "fmt"

func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

//Omitting a Parameter Name
func ommit_parameter_example(product string, price, _ float64) {
	taxAmount := price * 0.5
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

//Defining Varadic Parameters
func varadic_parameter_example(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

//Dealing with No Arguments for a Variadic Parameter
func varadic_no_argument_example(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func main() {

}

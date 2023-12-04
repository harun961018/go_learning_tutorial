package main

import "fmt"

func main() {
	kayakPrice := 275.00
	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}

	if kayakPrice > 100 && kayakPrice < 500 {
		fmt.Println("Price is greater than 100 and less than 500")
	}
	fmt.Println("Price:", kayakPrice)
}

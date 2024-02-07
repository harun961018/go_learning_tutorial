package main

import "fmt"

type calcFunc func(float64) float64

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	} else {
		return calcWithoutTax
	}
}

func function_comparisons_example() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

func function_argument_example() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}
}

func function_result_example() {
	products := map[string]float64{
		"Kayak":      123,
		"Lifejacket": 48.25,
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

func main() {
	// function_argument_example()
	function_result_example()
}

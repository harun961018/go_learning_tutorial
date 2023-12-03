package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

// This function writes a string using the fmt.Println function
func PrintHello() {
	fmt.Println("Hello, Go")
}

// This function writes a number using the fmt.Println function
func PrintNumber(number int) {
	fmt.Println(number)
}

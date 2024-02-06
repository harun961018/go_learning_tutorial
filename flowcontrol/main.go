package main

import (
	"fmt"
	"strconv"
)

//Using if Statements

func if_example() {
	kayakPrice := 275.00
	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}

	if kayakPrice > 100 && kayakPrice < 500 {
		fmt.Println("Price is greater than 100 and less than 500")
	}
	fmt.Println("Price:", kayakPrice)
}

//Using an Initialization Statement with an if Statement
func initialization_if_example() {
	priceString := "275"
	if kayakPrice, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayakPrice)
	} else {
		fmt.Println("Error:", err)
	}
}

//basic loop
func basic_loop_example() {
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}
}

//Incorporating Condition into the Loop
func incorporating_condition_loop() {
	counter := 0
	for counter < 3 {
		fmt.Println("Counter:", counter)
		counter++
	}
}

// Using Initialization and Completion Statements
func intialization_completion_loop() {
	for counter := 0; counter <= 3; counter++ {
		fmt.Println("Counter:", counter)
	}
}

//Using basic switch statements
func basic_switch_statements() {
	product := "Kayak"

	for index, character := range product {
		switch character {
		case 'K', 'k':
			fmt.Println("K at position", index)
			fallthrough
		case 'y':
			fmt.Println("y at position", index)
			break
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}
}

//Using Label Statements
func label_example() {
	counter := 0
target:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto target
	}
}

func main() {
	// basic_loop_example()
	// incorporating_condition_loop()
	// intialization_completion_loop()
	label_example()
}

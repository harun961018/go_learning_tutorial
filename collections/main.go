package main

import (
	"fmt"
	// "sort"
)

//basic array
func basic_array() {
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
}

//short syntax
func short_syntax_array_example() {
	names := [3]string{"Kayak", "top", "jacket"}
	fmt.Println(names)
}

//Comparing array
func comparing_array_example() {
	names := [3]string{"bottom", "top", "center"}
	otherNames := [3]string{"bottom", "top", "center"}
	same := names == otherNames
	fmt.Println(same)
}

//Enumerating Array
func enumerating_array_example() {
	names := [3]string{"bottom", "top", "center"}

	for index, name := range names {
		fmt.Println("Index:", index, "Value", name)
	}
}

//Basic slice
func basic_slice_example() {
	names := make([]string, 3)
	otherNames := []string{"a", "b", "c"}
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	otherNames = append(otherNames, "hat", "hg")
	anotherNames := append(names, "hat", "gloves")
	appendNames := append(names, anotherNames...)
	fmt.Println("names", names)
	fmt.Println("otherNames", otherNames)
	fmt.Println("anotherNames", anotherNames)
	fmt.Println("appendNames", appendNames)

}

//Creating Slices from Existing Arrays
func creat_slice_array_example() {
	product := [4]string{"a", "b", "c", "d"}
	someNames := product[1:3]
	allNames := product[:]
	someNames = append(someNames, "d")

	fmt.Println("someNames", someNames)
	fmt.Println("allNames", allNames)
}

//Basic Map
func map_syntax_example() {
	products := make(map[string]float64, 10)
	products["Kayak"] = 279
	products["Lifejacket"] = 48.95
	otherProdjcts := map[string]float64{
		"shirt":  100,
		"jacket": 150,
	}

	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}

	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	for _, key := range keys {
		fmt.Println("otherKey:", key, "value:", products[key])
	}
}

func main() {
	// basic_array()
	// comparing_array_example()
	// enumerating_array_example()
	basic_slice_example()
}

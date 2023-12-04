package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayfunc()
	fmt.Println("===============================")
	arrayMap()
}

func arrayfunc() {
	fmt.Println("Hello, Collections")
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	otherArray := &names
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("otherArray:", otherArray)
	for index, value := range names {
		fmt.Println("Index:", index, "Value:", value)
	}

	names1 := []string{"Kayak", "Lifejacket", "Paddle"}
	names1 = append(names1, "Hat", "Gloves")
	fmt.Println(names1)
}

func arrayMap() {
	products := map[string]float64{
		"kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}

	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}

	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}
}

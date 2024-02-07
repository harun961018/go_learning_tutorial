package main

import "fmt"

type Product struct {
	name, category string
	price float64
	*Supplier
}

type Supplier struct {
	name, city string
}

type StockLevel struct {
	Product
	Alternate Product
	count int
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price -10, supplier}
}

func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

// func struct_construct_example() {
// 	Products := [2]*Product{
// 		newProduct("Product 1", "Category 1", 100),
//         newProduct("Product 2", "Category 2", 200),
// 	}

// 	for _, p := range Products {
// 		fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price)
// 	}
// }

// func basic_struct_example() {
// 	stockItem := StockLevel{
// 		Product: Product{ "Kayak", "Sports", 1000.00 },
// 		Alternate: Product{"Lifejacket", "Sports", 2000.00},
// 		count: 100,
// 	}

// 	fmt.Println("Name: ", stockItem.Product.name)
// 	fmt.Println("All Name: ", stockItem.Alternate.name)
// }

// func compare_struct_example() {
// 	p1 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
// 	p2 := Product { name: "Kayak", category: "Watersports", price: 275.00 }
// 	p3 := Product { name: "Kayak", category: "Boats", price: 275.00 }
// 	fmt.Println("p1 == p2:", p1 == p2)
// 	fmt.Println("p1 == p3:", p1 == p3)
// }

// func struct_array_slice_map_example() {
// 	array := [1]StockLevel { 
// 		{
// 			Product: Product{ "Kayak", "Sports", 1000.00 },
//             Alternate: Product{"Lifejacket", "Sports", 2000.00},
//             count: 100,
// 		},
// 	}
// 	fmt.Println("Array:", array[0].Product.name)

// 	slice :=[]StockLevel {
// 		{
// 			Product: Product{ "Kayak", "Sports", 1000.00 },
//             Alternate: Product{"Lifejacket", "Sports", 2000.00},
//             count: 100,
// 		},
// 	}
// 	fmt.Println("Slice:", slice[0].Product.name)

// 	kvp := map[string]StockLevel {
// 		"Kayak": {
// 			Product: Product{ "Kayak", "Sports", 1000.00 },
//             Alternate: Product{"Lifejacket", "Sports", 2000.00},
//             count: 100,
// 		},
// 	}

// 	fmt.Println("Map:", kvp["Kayak"].Product.name)
// }

// func struct_pointer_example() {
// 	p1 := Product {
// 		name: "Kayak",
// 		category: "Watersports",
//         price: 275.00,	
// 	}

// 	p2 := &p1

// 	p1.name = "Lifejacket"
// 	fmt.Println("p1.name:", p1.name)
// 	fmt.Println("p2.name:", (*p2).name)

// }

func struct_pointer_example_2() {
	acme := &Supplier { "Acme Co", "New York" }
	products := [2]*Product {
		newProduct("Product 1", "Category 1", 100, acme),
        newProduct("Product 2", "Category 2", 200, acme),
	}
	for _, p := range products {
        fmt.Println("Name:", p.name, "Category:", "Supplier:", p.Supplier.name, p.Supplier.city)
    }
}

func struct_pointer_copy_example() {
	acme := &Supplier { "Acme Co", "New York" }
	p1 := newProduct("Kayak", "Sports", 1000.00, acme)
	p2 := copyProduct(p1)
	p1.name = "Lifejacket"
	p1.Supplier.name = "Boat co"

	for _, p := range []Product { *p1, p2 } {
		fmt.Println("Name:", p.name, "Category:", "Supplier:", p.Supplier.name, p.Supplier.city)
	}
}

func interface_sytax_example() {
	expenses := []Expense {
		Product{"Kayak", "Sports", 1000.00},
		Service{"Boat", "Sports
	}

}

func main() {
	// basic_struct_example()
	// compare_struct_example()
	// struct_array_slice_map_example()	
	// struct_pointer_example()
	// struct_construct_example()
	struct_pointer_example_2()
	struct_pointer_copy_example()
}
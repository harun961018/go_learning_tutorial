package main

import (
    "fmt"
)


type Productlist []Product

func newProduct(name string, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) printDetails() {
	fmt.Printf("Name: %s, Category: %s, Price: %.2f\n", p.Name, p.Category, p.calcTax(0.2, 100))
}

func (products *Productlist) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.Category] += totals[p.Category] + p.Price
	}
	return totals

}
func (product *Product) calcTax(rate, threshold float64) float64 {
	if (product.Price > threshold) {
		return product.Price + (product.Price * rate)
	}
	return product.Price
}

func method_alias_example() {
	products := Productlist {
		{"Kayak", "Sports", 1000.00},
		{"Lifejacket", "Sports", 2000.00},
		{"Soccer Ball", "Sports", 3000.00},
	}
	for category, total := range products.calcCategoryTotals() { 
		fmt.Println("Category:", category, "Total:", total)
	}
}

func method_syntax_example() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 279),
        newProduct("Lifejacket", "Watersports", 49.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}

	for _, p := range products {
        p.printDetails()
    }
}



func main() {
    // method_syntax_example()
	method_alias_example()
}



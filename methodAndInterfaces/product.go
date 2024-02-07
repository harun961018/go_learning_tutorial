package main

type Product struct {
    Name, Category string
    Price          float64
}

func (p Product) getName() string {
	return p.Name
}

func (p Product) getCost(_ bool) float64 {
	return p.Price
}
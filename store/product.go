package store

import "database/sql"

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price(taxRate float64) float64 {
	return p.price + (p.price * taxRate)
}

func queryDatabase(db *sql.DB) {
	rows, err := db.Query("SELECT * from Products")
	if err == nil {
		for rows.Next() {
			var id, category int
			var name int
			var price float64
			scanErr := rows.Scan(&id, &name, &category, &price)
			if scanErr == nil {
				Printfln("Row: %v %v %v %v", id, name, category, price)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
}

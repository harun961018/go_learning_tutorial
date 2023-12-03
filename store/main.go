package store

import (
	"database/sql"
)

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

func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if err == nil {
		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if err != nil || updatedFailed {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

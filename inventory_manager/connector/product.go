package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteProduct(db *sql.DB, id string, brand string, name string, cost_price float64, selling_price float64) (int64, error) {
	result, err := db.Exec(`INSERT INTO product (id, brand, name, cost_price, selling_price) VALUES (?, ?, ?, ?, ?)`, id, brand, name, cost_price, selling_price)
	if err != nil {
		log.Println("Error inserting product:", err)
	}
	return result.LastInsertId()
}

func FetchProduct(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

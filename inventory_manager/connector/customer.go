package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteCustomer(db *sql.DB, id string, name string, contact string, email string) (int64, error) {
	result, err := db.Exec(`INSERT INTO customer (id, name, contact, email) VALUES (?, ?, ?, ?)`, id, name, contact, email)
	if err != nil {
		log.Println("Error inserting customer:", err)
	}
	return result.LastInsertId()
}

func FetchCustomer(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

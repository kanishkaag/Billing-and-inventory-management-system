package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteSupplier(db *sql.DB, id string, name string, contact string, address string, email string) (int64, error) {
	result, err := db.Exec(`INSERT INTO supplier (id, name, contact, address, email) VALUES (?, ?, ?, ?, ?)`, id, name, contact, address, email)
	if err != nil {
		log.Println("Error inserting supplier:", err)
	}
	return result.LastInsertId()
}

func FetchSupplier(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

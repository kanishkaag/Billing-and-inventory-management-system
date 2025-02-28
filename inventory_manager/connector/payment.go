package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WritePayment(db *sql.DB, id string, bill_id string, status string, medium string) (int64, error) {
	result, err := db.Exec(`INSERT INTO payment (id, bill_id, status, medium) VALUES (?, ?, ?, ?)`, id, bill_id, status, medium)
	if err != nil {
		log.Println("Error inserting payment:", err)
	}
	return result.LastInsertId()
}

func FetchPayment(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

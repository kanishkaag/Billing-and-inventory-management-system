package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteOutlet(db *sql.DB, id string, shop_id string, address string, contact string, email string) (int64, error) {
	result, err := db.Exec(`INSERT INTO outlet (id, shop_id, address, contact, email) VALUES (?, ?, ?, ?, ?)`, id, shop_id, address, contact, email)
	if err != nil {
		log.Println("Error inserting outlet:", err)
	}
	return result.LastInsertId()
}

func FetchOutlet(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteShop(db *sql.DB, id string, name string) (int64, error) {
	result, err := db.Exec(`INSERT INTO shop (id, name) VALUES (?, ?)`, id, name)
	if err != nil {
		log.Println("Error inserting shop:", err)
	}
	return result.LastInsertId()
}

func FetchShop(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

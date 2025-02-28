package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteStaff(db *sql.DB, id string, outlet_id string, name string, user string, password string, contact string) (int64, error) {
	result, err := db.Exec(`INSERT INTO staff (id, outlet_id, name, user, password, contact) VALUES (?, ?, ?, ?, ?, ?)`, id, outlet_id, name, user, password, contact)
	if err != nil {
		log.Println("Error inserting staff:", err)
	}
	return result.LastInsertId()
}

func FetchStaff(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

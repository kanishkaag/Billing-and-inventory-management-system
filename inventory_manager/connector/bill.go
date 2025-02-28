package connector

import (
	"database/sql"
	"log"
	"newBackend/models"
)

func WriteBill(db *sql.DB, customer_id int, shopkeeper_id int, total_amount float64, payment_status string, payment_mode string) (int64, error) {
	result, err := db.Exec(`insert into bill (customer_id, shopkeeper_id, total_amount, payment_status) values(?,?,?,?)`, customer_id, shopkeeper_id, total_amount, payment_status)
	if err != nil {
		log.Println("Error generating bill ", err)
	}
	return result.LastInsertId()
}

func FetchBill(query string, result any, args ...any) error {
	err := models.DBClient.ReadMany(result, query, args...)
	return err
}

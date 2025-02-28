package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySqlClient struct {
	db *sqlx.DB
}

func getMySQLClient(config connectionConfig) (*MySqlClient, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.username,
		config.password,
		config.host,
		config.port,
		config.database,
	)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	client := &MySqlClient{db: db}
	log.Println("Connected to MySQL database!")
	return client, nil
}

// ReadOne fetches a single record by a specified query and destination
func (client *MySqlClient) ReadOne(dest interface{}, query string, args ...interface{}) error {
	return client.db.Get(dest, query, args...)
}

// ReadMany fetches multiple records by a specified query and destination
func (client *MySqlClient) ReadMany(dest interface{}, query string, args ...interface{}) error {
	return client.db.Select(dest, query, args...)
}

// WriteOne executes an insert query for a single record
func (client *MySqlClient) WriteOne(query string, args ...interface{}) (sql.Result, error) {
	return client.db.Exec(query, args...)
}

// WriteMany executes an insert query for multiple records in a transaction
func (client *MySqlClient) WriteMany(query string, records []interface{}) error {
	tx, err := client.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		if _, err := stmt.Exec(record); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteOne deletes a single record by a specified query
func (client *MySqlClient) DeleteOne(query string, args ...interface{}) (sql.Result, error) {
	return client.db.Exec(query, args...)
}

// DeleteMany deletes multiple records by a specified query
func (client *MySqlClient) DeleteMany(query string, args ...interface{}) (sql.Result, error) {
	return client.db.Exec(query, args...)
}

// UpdateOne updates a single record by a specified query
func (client *MySqlClient) UpdateOne(query string, args ...interface{}) (sql.Result, error) {
	return client.db.Exec(query, args...)
}

// UpdateMany updates multiple records by a specified query
func (client *MySqlClient) UpdateMany(query string, records []interface{}) error {
	tx, err := client.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		if _, err := stmt.Exec(record); err != nil {
			return err
		}
	}

	return tx.Commit()
}

package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	// postgres driver

	_ "github.com/lib/pq"
)

type SQLDatabase interface {
	CreateDataDump(dd BTDataDump) error
}

type BTDataDump struct {
	ID         int      `json:"id,omitempty"`
	Sensor     string   `json:"sensor,omitempty"`
	DataValues []string `json:"data_values,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
}

type Db struct {
	*sql.DB
}

func NewDb() (*Db, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("BT_DATA_SERVER_DB_HOST"),
		os.Getenv("BT_DATA_SERVER_DB_PORT"),
		os.Getenv("BT_DATA_SERVER_DB_USER"),
		os.Getenv("BT_DATA_SERVER_DB_PASSWORD"),
		os.Getenv("BT_DATA_SERVER_DB_NAME"),
		os.Getenv("BT_DATA_SERVER_SSL_MODE"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func (db *Db) CreateDataDump(dd BTDataDump) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dataValues, err := json.Marshal(dd.DataValues)
	if err != nil {
		return err
	}

	query := "INSERT INTO data_dumps (sensor, data_values) VALUES ($1, $2);"
	if _, err := tx.Exec(query, dd.Sensor, dataValues); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

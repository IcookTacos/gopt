package storage

import (
	"database/sql"
	"log"
)

func Insert(key string, value string) error{
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
    return err
	}

	defer db.Close()

  _, err := db.Exec(
    `INSERT OR REPLACE INTO key_value_pairs (key, value)
    VALUES (?, ?)`,
    key, value)

  return err
}

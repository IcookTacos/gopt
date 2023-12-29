package storage

import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/mattn/go-sqlite3"
)

func Store(key string, value string) error{
  db, err := sql.Open("sqlite3", "storage/database.db")

  if err!=nil {
    log.Println("Error when opening databse")
    return err
  }

	defer db.Close()

  _, err = db.Exec("INSERT OR REPLACE INTO key_value_pairs (key, value) VALUES (?, ?)", key, value)

  if err!=nil {
    log.Println("Error when doing INSERT", err)
    return err
  }

  log_message := fmt.Sprintf("Inserting %s : %s", key, value)
  log.Println(log_message)

  return err
}

func List(key string) (error, string) {
  db, err := sql.Open("sqlite3", "storage/database.db")

  if err!=nil {
    log.Println("Error when interfacing with databse")
    return err, ""
  }

	defer db.Close()

  rows, err := db.Query("SELECT value FROM key_value_pairs WHERE key == (?)", key)

  if err != nil {
		log.Println("Error executing SELECT query:", err)
		return err, ""
	}

	defer rows.Close()

	var value string
  for rows.Next() {
		if err := rows.Scan(&value); err != nil {
			log.Println("Error scanning rows:", err)
			return err, ""
		}
	}
    return err, value
}

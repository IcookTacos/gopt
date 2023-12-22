package storage

import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/mattn/go-sqlite3"
)

func Insert(key string, value string) error{
  db, err := sql.Open("sqlite3", "storage/database.db")

  if(err!=nil){
    log.Println("Error when interfacing with databse")
    return err
  }

	defer db.Close()

  _, err = db.Exec("INSERT OR REPLACE INTO key_value_pairs (key, value) VALUES (?, ?)", key, value)

  if(err!=nil){
    log.Println("Error when doing INSERT", err)
    return err
  }

  log_message := fmt.Sprintf("Inserting %s : %s", key, value)
  log.Println(log_message)

  return err
}

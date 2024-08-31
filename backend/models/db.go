package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := os.Getenv("username") + `:` + os.Getenv("password") + `@tcp(34.89.177.188:3306)/task_manager`

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	fmt.Println("successfully connected to database")
	DB = db
}

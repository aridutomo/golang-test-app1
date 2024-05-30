package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/crudapp")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping the database:", err)
		return
	}

	fmt.Println("Connected to the database!")

	// Close the database connection when done
	defer db.Close()

	// Use the db variable to perform database operations
	// ...

	// $PLACEHOLDER$
}

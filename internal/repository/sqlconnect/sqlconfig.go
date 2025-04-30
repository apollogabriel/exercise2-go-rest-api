package sqlconnect

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(dbname string) (*sql.DB, error) {
	fmt.Println("Connecting to database...")
	connectionString := fmt.Sprintf("root:root@tcp(localhost:3306)/%s", dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		//panic(err)
		return nil, err
	}
	fmt.Println("Connected to database...")
	return db, nil
}

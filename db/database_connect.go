package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase() (*sql.DB, error) {
	// serverName := "localhost:3306"
	// user := "albert"
	// password := "CDLcdl"
	// dbName := "petStore"

	serverName := "localhost:3307"
	user := "root"
	password := "Solutions!"
	dbName := "petStore"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic("Database conncetion failed!", err)
		return nil, err
	}

	return db, nil
}

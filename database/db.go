package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func DB() (db *sql.DB) {
	driver := "mysql"
	db, err := sql.Open(driver, "doadmin:AVNS_326MIfcmreLTEOq1Si3@tcp(db-mysql-nyc1-41608-do-user-14594785-0.b.db.ondigitalocean.com:25060)/defaultdb")
	if err != nil {
		log.Fatal(err) // Return the error if the database connection fails
	}
	return db
}

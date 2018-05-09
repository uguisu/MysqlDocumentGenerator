package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
 * Get database connection
 */
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.11.120:3306)/himysql?charset=utf8");

	if err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		return db, nil
	}
	
}
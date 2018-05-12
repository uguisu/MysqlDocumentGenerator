package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Start")

	// init connection information
	mysqlConnectInfo := ConnectionInformation{
		"root", "root", "himysql", "3306", "192.168.11.120",
	}

	// Get database
	db, err := getDB(mysqlConnectInfo)
	checkErr(err)

	// make sure db connection will be closed
	defer db.Close()

	// Execute Sql statement
	rows, err := db.Query(SQL)
	checkErr(err)

	// Declare table variable
	var tabelInfoRow TabelInfo
	tableRecordCollection := make([]TabelInfo, 0)

	// Fetch data
	for rows.Next() {

		err = rows.Scan(
			&tabelInfoRow.tabel_name,
			&tabelInfoRow.table_comment,
			&tabelInfoRow.column_name,
			&tabelInfoRow.collation_name,
			&tabelInfoRow.data_type,
			&tabelInfoRow.character_maximum_length,
			&tabelInfoRow.column_key,
			&tabelInfoRow.is_nullable,
			&tabelInfoRow.column_comment,
		)
		checkErr(err)

		log.Printf("tabel_name= %s, table_comment= %s, column_name=%s, collation_name= %s, data_type= %s, character_maximum_length= %s, column_key= %s, is_nullable= %s, column_comment= %s \n",
			tabelInfoRow.tabel_name,
			tabelInfoRow.table_comment,
			tabelInfoRow.column_name,
			tabelInfoRow.collation_name,
			tabelInfoRow.data_type,
			tabelInfoRow.character_maximum_length,
			tabelInfoRow.column_key,
			tabelInfoRow.is_nullable,
			tabelInfoRow.column_comment,
		)

		tableRecordCollection = append(tableRecordCollection, tabelInfoRow)
	}

	log.Printf("total = %d", len(tableRecordCollection))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * Get database connection
 */
func getDB(con ConnectionInformation) (*sql.DB, error) {

	connectString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		con.user,
		con.pwd,
		con.serverName,
		con.port,
		con.schema,
	)

	db, err := sql.Open("mysql", connectString)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		return db, nil
	}
}

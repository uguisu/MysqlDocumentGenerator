package main

import (
	"github.com/uguisu/connection"
	"log"
)

func main() {
	log.Println("Start")

	// Get database
	db, err := connection.GetDB()
	checkErr(err)

	// make sure db connection will be closed
	defer db.Close()

	// Execute Sql statement
	rows, err := db.Query(SQL)
	checkErr(err)

	// Declare table variable
	var tabelInfoRow tabelInfo
	tableRecordCollection := make([]tabelInfo, 0)

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
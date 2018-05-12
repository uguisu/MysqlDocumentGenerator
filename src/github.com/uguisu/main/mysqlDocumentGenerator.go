package main

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/uguisu/config"
	connectionInfo "github.com/uguisu/connectionInfo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Start")

	// Load config
	configFinished := config.LoadConfig()
	if <-configFinished {
		log.Println("Load config finished.")
	} else {
		log.Fatal("Unknow excetion whien loading config")
	}

	// Get database
	db, err := getDB(config.GetConnectionInfoObject())
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

		// log.Printf("tabel_name= %s, table_comment= %s, column_name=%s, collation_name= %s, data_type= %s, character_maximum_length= %s, column_key= %s, is_nullable= %s, column_comment= %s \n",
		// 	tabelInfoRow.tabel_name,
		// 	tabelInfoRow.table_comment,
		// 	tabelInfoRow.column_name,
		// 	tabelInfoRow.collation_name,
		// 	tabelInfoRow.data_type,
		// 	tabelInfoRow.character_maximum_length,
		// 	tabelInfoRow.column_key,
		// 	tabelInfoRow.is_nullable,
		// 	tabelInfoRow.column_comment,
		// )

		tableRecordCollection = append(tableRecordCollection, tabelInfoRow)
	}

	log.Printf("total = %d", len(tableRecordCollection))

	var lastTableName string = tableRecordCollection[0].tabel_name
	var lastStartIndex = 0
	tableMap := make(map[string][]TabelInfo, 0)
	for i, val := range tableRecordCollection {
		if lastTableName != val.tabel_name {

			log.Printf("find = %s", val.tabel_name)

			// tabel changed
			tableMap[lastTableName] = tableRecordCollection[lastStartIndex:i]
			lastStartIndex = i
			lastTableName = val.tabel_name
		}
	}
	// tabel changed
	tableMap[lastTableName] = tableRecordCollection[lastStartIndex:len(tableRecordCollection)]

	for k, v := range tableMap {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * Get database connection
 */
func getDB(in *connectionInfo.ConnectionInfo) (*sql.DB, error) {

	db, err := sql.Open("mysql", in.GetConnectionString())

	if err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		return db, nil
	}
}

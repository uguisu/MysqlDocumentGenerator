package main

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/uguisu/config"
	dbconnect "github.com/uguisu/dbconnect"

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

	// Read and transfer data
	tableMap := transferRowsToMap(rows)

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
func getDB(in *dbconnect.ConnectionInfo) (*sql.DB, error) {

	db, err := sql.Open("mysql", in.GetConnectionString())

	if err != nil {
		log.Fatalln(err)
		return nil, err
	} else {
		return db, nil
	}
}

/**
 * Read and transfer data
 */
func transferRowsToMap(rows *sql.Rows) map[string][]TabelInfo {
	// Declare table variable
	var tabelInfoRow TabelInfo
	tableRecordCollection := make([]TabelInfo, 0)

	// Fetch data
	for rows.Next() {

		err := rows.Scan(
			&tabelInfoRow.tabelName,
			&tabelInfoRow.tableComment,
			&tabelInfoRow.columnName,
			&tabelInfoRow.collationName,
			&tabelInfoRow.dataType,
			&tabelInfoRow.characterMaximumLength,
			&tabelInfoRow.columnKey,
			&tabelInfoRow.isNullable,
			&tabelInfoRow.columnComment,
		)
		checkErr(err)

		tableRecordCollection = append(tableRecordCollection, tabelInfoRow)
	}

	log.Printf("total = %d", len(tableRecordCollection))

	var lastTableName string = tableRecordCollection[0].tabelName
	var lastStartIndex = 0
	tableMap := make(map[string][]TabelInfo, 0)
	for i, val := range tableRecordCollection {
		if lastTableName != val.tabelName {

			log.Printf("find = %s", val.tabelName)

			// tabel changed
			tableMap[lastTableName] = tableRecordCollection[lastStartIndex:i]
			lastStartIndex = i
			lastTableName = val.tabelName
		}
	}
	// tabel changed
	tableMap[lastTableName] = tableRecordCollection[lastStartIndex:len(tableRecordCollection)]

	return tableMap
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	rows, err := db.Query(fmt.Sprintf(SQL, config.GetConnectionInfoObject().Schema))
	checkErr(err)

	// Read and transfer data
	tableMap := transferRowsToMap(rows)

	writeToMd(tableMap)
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
	log.Printf("find = %s", lastTableName)

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

/**
 * Write to file in markdown format
 */
func writeToMd(tableMap map[string][]TabelInfo) {
	var rowStr = ""
	var pk = ""
	var err error = nil

	// create output file
	f, err := os.Create(OutputFileName)
	checkErr(err)

	defer f.Close()

	for k, v := range tableMap {
		// Title
		title := fmt.Sprintf("\n## %s\n", k)
		// Comments
		commetns := fmt.Sprintf("_%s_\n\n", v[1].tableComment)

		// output columns info
		for _, rloop := range v {

			if rloop.columnKey == "PRI" {
				pk = fmt.Sprintf("`%s`", rloop.columnName)
			} else {
				pk = rloop.columnName
			}

			rowStr = fmt.Sprintf("%s| %s | %s | %s | %s | %s |\n",
				rowStr,
				pk,
				rloop.dataType,
				rloop.characterMaximumLength,
				rloop.isNullable,
				rloop.columnComment,
			)
		}

		_, err = f.Write([]byte(title))
		_, err = f.Write([]byte(commetns))
		_, err = f.Write([]byte(TableHeader1))
		_, err = f.Write([]byte(TableHeader2))

		_, err = f.Write([]byte(rowStr))

		rowStr = ""
	}

	f.Sync()
}

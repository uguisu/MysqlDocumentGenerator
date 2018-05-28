package main

// SQL statement
const SQL string = "SELECT " +
	"    t.TABLE_NAME, " +
	"    t.TABLE_COMMENT, " +
	"    c.COLUMN_NAME, " +
	"    ifnull(c.COLLATION_NAME,''), " +
	// "    c.COLLATION_NAME, " +
	"    c.DATA_TYPE, " +
	"    ifnull(c.CHARACTER_MAXIMUM_LENGTH, ''), " +
	// "    c.CHARACTER_MAXIMUM_LENGTH, " +
	"    c.COLUMN_KEY, " +
	"    c.IS_NULLABLE, " +
	"    c.COLUMN_COMMENT " +
	"FROM " +
	"    information_schema.tables t " +
	"    inner join " +
	"    information_schema.columns c " +
	"    on " +
	"    t.table_schema = c.table_schema " +
	"    AND t.table_name = c.table_name " +
	"WHERE " +
	"    t.table_type = 'BASE TABLE' " +
	"    AND t.table_schema = 'himysql' " +
	"ORDER BY " +
	"    t.table_name, c.ordinal_position"

// Output file name
const OutputFileName string = "Database.md"

const TableHeader1 string = "| Column Name | Type | Length | Nullable | Column Comment |\n"
const TableHeader2 string = "| --------------- | --------------- | --------------- | --------------- | --------------- |\n"

// Table record structor
type TabelInfo struct {
	tabelName              string
	tableComment           string
	columnName             string
	collationName          string
	dataType               string
	characterMaximumLength string
	columnKey              string
	isNullable             string
	columnComment          string
}

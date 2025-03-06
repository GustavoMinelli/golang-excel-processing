package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {

	var err error

	// Connect to the database
	db, err = sql.Open("postgres", "user=postgres dbname=todo password=postgres sslmode=disable")

	// Check for errors
	if err != nil {
		panic(err)
	}

	// Print a success message
	fmt.Println("Successfully connected to the database!")

}

func GetData(query string) ([]map[string]any, error) {

	rows, err := db.Query("SELECT * FROM table WHERE table.relation_id = 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}

	for rows.Next() {

		columnsData := make([]interface{}, len(columns))
		columnsPointers := make([]interface{}, len(columns))

		for i := range columnsData {
			columnsPointers[i] = &columnsData[i]
		}

		if err := rows.Scan(columnsPointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})

		for i, colName := range columns {
			val := columnsPointers[i].(*interface{})
			rowMap[colName] = *val
		}

		results = append(results, rowMap)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func CloseConnection() {

	if db != nil {
		db.Close()
		fmt.Println("Successfully closed the database connection!")
	}

}

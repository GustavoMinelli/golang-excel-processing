package database

import (
	"database/sql"
	"fmt"

	"github.com/GustavoMinelli/golang-excel-processing/internal"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() {

	var err error
	config, err := internal.GetConfig()
	database := config.Database

	if err != nil {
		panic(err)
	}

	db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", database.User, database.Name, database.Password, database.Sslmode))

	// Check for errors
	if err != nil {
		panic(err)
	}

}

func GetData(query string) ([]map[string]any, error) {

	rows, err := db.Query(query)

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

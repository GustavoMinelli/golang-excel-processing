package main

import (
	"fmt"

	"github.com/GustavoMinelli/golang-excel-processing/internal/database"
	"github.com/GustavoMinelli/golang-excel-processing/internal/excel"
)

// Main function
func main() {

	database.Connect()
	data, err := database.GetData()
	defer database.CloseConnection()

	if err != nil {
		fmt.Println(err)
		return
	}

	excel.ExportData(data)

}

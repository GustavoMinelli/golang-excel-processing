package excel

import (
	"fmt"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func ExportData(data []map[string]any, path string) {

	excel := excelize.NewFile()

	// Create a new sheet.
	index, err := excel.NewSheet("Date")

	if err != nil {
		fmt.Println(err)
		return
	}

	excel.SetCellValue("Date", "A1", "Date")
	excel.SetCellValue("Date", "B1", "Value")

	for i := 0; i < len(data); i++ {

		date := data[i]["date"]
		value := data[i]["value"]

		// Set value of a cell.
		excel.SetCellValue("Date", fmt.Sprintf("A%d", i+2), date)
		excel.SetCellValue("Date", fmt.Sprintf("B%d", i+2), value)
	}

	// Set active sheet of the workbook.
	excel.SetActiveSheet(index)

	outputPath := filepath.Join(path, "Test.xlsx")

	// Save xlsx file by the given path.
	if err := excel.SaveAs(outputPath); err != nil {
		fmt.Println(err)
		return
	}

	println("Excel file created successfully!")

}

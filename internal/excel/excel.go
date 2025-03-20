package excel

import (
	"fmt"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func ExportData(data []map[string]any, path string, fileName string, rowTitle []string) error {

	excel := excelize.NewFile()

	// Create a new sheet.
	index, err := excel.NewSheet("Date")

	if err != nil {
		return err
	}

	excel.SetCellValue("Date", "A1", "Date")
	excel.SetCellValue("Date", "B1", "Value")

	for i, row := range data {
		date := row["date"]
		value := row["value"]

		for j, title := range rowTitle {
			excel.SetCellValue("Date", fmt.Sprintf("%s%d", string(rune(65+j)), 1), title)
		}

		excel.SetCellValue("Date", fmt.Sprintf("A%d", i+2), date)
		excel.SetCellValue("Date", fmt.Sprintf("B%d", i+2), value)
	}

	// Set active sheet of the workbook.
	excel.SetActiveSheet(index)

	outputPath := filepath.Join(path, "Test.xlsx")

	// Save xlsx file by the given path.
	if err := excel.SaveAs(outputPath); err != nil {
		return err
	}

	return nil
}

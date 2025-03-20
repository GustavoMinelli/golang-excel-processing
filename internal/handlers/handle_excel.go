package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoMinelli/golang-excel-processing/internal/database"
	"github.com/GustavoMinelli/golang-excel-processing/internal/excel"
	"github.com/aws/aws-lambda-go/events"
)

// Requirements to build the excel file (Sent by the client)
type Excel struct {
	FileName string   `json:"file_name"`
	RowTitle []string `json:"row_title"`
	Query    string   `json:"query"`
}

// Handle Excel
func HandleExcel(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	var excelRequest Excel

	// Validate the request
	err := validator.Validate(req, &excelRequest)

	err = json.Unmarshal([]byte(req.Body), &excelRequest)

	if err != nil {
		return HandleError(err)
	}

	query := req.QueryStringParameters["query"]

	database.Connect()
	data, err := database.GetData(query)
	defer database.CloseConnection()

	if err != nil {
		return HandleError(err)
	}

	err = excel.ExportData(data, "output", excelRequest.FileName, excelRequest.RowTitle)

	if err != nil {
		return HandleError(err)
	}

	// Create the response
	response := Response{
		Message: "Excel file processed successfully",
		Path:    "test",
	}

	responseBody, err := json.Marshal(response)

	if err != nil {
		return HandleError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(responseBody),
	}
}

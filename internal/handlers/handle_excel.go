package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoMinelli/golang-excel-processing/internal/database"
	"github.com/GustavoMinelli/golang-excel-processing/internal/excel"
	"github.com/aws/aws-lambda-go/events"
)

// Handle Excel
func HandleExcel(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	query := req.QueryStringParameters["query"]

	database.Connect()
	data, err := database.GetData(query)

	defer database.CloseConnection()

	if err != nil {
		HandleError(err)
	}

	excel.ExportData(data, "output")

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

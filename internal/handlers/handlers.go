package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GustavoMinelli/golang-excel-processing/internal/database"
	"github.com/GustavoMinelli/golang-excel-processing/internal/excel"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func HandleExcel(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	query := req.QueryStringParameters["query"]

												
	database.Connect()
	data, err := database.GetData(query)

	defer database.CloseConnection()

	if err != nil {
		handleError(err)
	}

	excel.ExportData(data)

	// Create the response
	response := Response{
		Message: "Excel file processed successfully",
		Path:    "test",
	}

	responseBody, err := json.Marshal(response)

	if err != nil {
		return handleError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(responseBody),
	}
}

func handleError(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	}
}

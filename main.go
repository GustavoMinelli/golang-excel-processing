package main

import (
	"net/http"

	"github.com/GustavoMinelli/golang-excel-processing/internal/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Main function
func main() {
	lambda.Start(router)
}

// Router
func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.Path == "/excel" {

		if req.HTTPMethod == "POST" {
			return handlers.HandleExcel(req), nil
		}

	}

	//Method not fonud
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}, nil

}

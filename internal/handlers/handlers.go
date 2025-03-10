package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func HandleError(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	}
}

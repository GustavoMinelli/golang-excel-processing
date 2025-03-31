package main

import (
	"fmt"
	"net/http"

	"github.com/GustavoMinelli/golang-excel-processing/internal/handlers"
	"github.com/aws/aws-lambda-go/events"
)

// Main function
func main() {
	// lambda.Start(router)
	sandBox()
}

// Router
func router(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {

	if req.Path == "/excel" {

		if req.HTTPMethod == "POST" {
			return handlers.HandleExcel(req)
		}

	}

	//Method not fonud
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       http.StatusText(http.StatusNotFound),
	}

}

func sandBox() {

	// Exemplo de evento para testar o HandleExcel
	event := events.APIGatewayProxyRequest{
		Path:       "/excel",
		HTTPMethod: "POST",
		Body: `{
            "query": "SELECT * FROM USERS",
            "row_title": ["Nome do índice", "Valor", "Data"],
            "file_name": "example_file.xlsx"
        }`,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	// Chama o handler e captura a resposta
	response := router(event)

	// Imprime a resposta para verificar o resultado
	fmt.Printf("Status Code: %d\n", response.StatusCode)
	fmt.Printf("Body: %s\n", response.Body)
	fmt.Printf("Headers: %v\n", response.Headers)
}

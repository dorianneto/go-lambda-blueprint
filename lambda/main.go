package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type responseBody struct {
	Message string `json:"message"`
}

func payload(message string) string {
	payload, _ := json.Marshal(responseBody{Message: message})

	return string(payload)
}

func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return events.APIGatewayProxyResponse{
			Body:       payload("Hello World!"),
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, nil
	})
}

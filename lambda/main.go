package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type input struct {
	Link string `json:"link"`
}

type responseBody struct {
	Message string `json:"message"`
}

func payload(message string, err error) string {
	payload, _ := json.Marshal(responseBody{Message: fmt.Sprintf("%s (%s)", message, err)})

	return string(payload)
}

func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var input input

		body := request.Body

		err := json.Unmarshal([]byte(body), &input)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       payload("Could not process input", nil),
			}, err
		}

		return events.APIGatewayProxyResponse{
			Body:       payload("sadasdad", nil),
			StatusCode: http.StatusOK,
		}, nil
	})
}

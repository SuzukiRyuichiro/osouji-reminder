// package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// )

// type App struct {
// 	id string
// }

// func (a *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	responseBody := map[string]string{
// 		"message": "never gonna give you up",
// 	}

// 	responseJSON, err := json.Marshal(responseBody)

// 	if err != nil {
// 		return events.APIGatewayProxyResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Headers:    map[string]string{"Content-Type": "application/json"},
// 			Body:       `{"error":"internal server error"}`,
// 		}, nil
// 	}

// 	response := events.APIGatewayProxyResponse{
// 		Body:       string(responseJSON),
// 		StatusCode: http.StatusOK,
// 		Headers: map[string]string{
// 			"Content-Type":                     "text/plain",
// 			"Access-Control-Allow-Origin":      "*",
// 			"Access-Control-Allow-Headers":     "Content-Type",
// 			"Access-Control-Allow-Methods":     "OPTIONS, POST, GET, PUT, DELETE",
// 			"Access-Control-Allow-Credentials": "true",
// 		},
// 	}

// 	return response, nil
// }

// func newApp(id string) *App {
// 	return &App{
// 		id: id,
// 	}
// }

// func main() {
// 	id := "someId"
// 	app := newApp(id)
// 	lambda.Start(app.Handler)
// }

package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (string, error) {
	return "Hello, World!", nil
}

func main() {
	lambda.Start(handler)
}

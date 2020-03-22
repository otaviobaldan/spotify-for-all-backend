package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func main() {
	lambda.Start(router)
}

func handleCreateUser(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var user User

	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return BadRequest(err)
	}

	userNew, err := CreateUser(user)
	if err != nil || userNew == nil {
		return BadRequest(err)
	}
	response, err := json.Marshal(userNew)
	if err != nil {
		return BadRequest(err)
	}
	return Created(string(response))
}

func handleGetUsers(_ events.APIGatewayProxyRequest) events.APIGatewayProxyResponse  {
	users, err := GetUsers()
	if err != nil {
		return BadRequest(err)
	}

	js, err := json.Marshal(users)
	if err != nil {
		BadRequest(err)
	}

	return Ok(string(js))
}

func router(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	switch request.Path {
	case "/users":
		{
			if request.HTTPMethod == http.MethodPost {
				return handleCreateUser(request)
			}
			if request.HTTPMethod == http.MethodGet {
				return handleGetUsers(request)
			}
		}
	default:
		return MethodNotAllowed(request.HTTPMethod)
	}

	return MethodNotAllowed(request.HTTPMethod)
}
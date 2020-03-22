package controller

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/otaviobaldan/spotify-for-all-backend/models"
	"github.com/otaviobaldan/spotify-for-all-backend/responses"
	"github.com/otaviobaldan/spotify-for-all-backend/service"
)

func HandleCreateUser(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var user models.User

	if err := json.Unmarshal([]byte(request.Body), &user); err != nil {
		return responses.BadRequest(err)
	}

	userNew, err := service.CreateUser(user)
	if err != nil || userNew == nil {
		return responses.BadRequest(err)
	}
	response, err := json.Marshal(userNew)
	if err != nil {
		return responses.BadRequest(err)
	}
	return responses.Created(string(response))
}

func HandleGetUsers(_ events.APIGatewayProxyRequest) events.APIGatewayProxyResponse  {
	users, err := service.GetUsers()
	if err != nil {
		return responses.BadRequest(err)
	}

	js, err := json.Marshal(users)
	if err != nil {
		responses.BadRequest(err)
	}

	return responses.Ok(string(js))
}


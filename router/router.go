package router

import (
	"github.com/otaviobaldan/spotify-for-all-backend/controller"
	"github.com/otaviobaldan/spotify-for-all-backend/responses"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Router (request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	switch request.Path {
	case "/users":
		{
			if request.HTTPMethod == http.MethodPost {
				return controller.HandleCreateUser(request)
			}
			if request.HTTPMethod == http.MethodGet {
				return controller.HandleGetUsers(request)
			}
		}
	default:
		return responses.MethodNotAllowed(request.HTTPMethod)
	}

	return responses.MethodNotAllowed(request.HTTPMethod)
}

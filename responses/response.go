package responses

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func BadRequest(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       err.Error(),
	}
}

func MethodNotAllowed(httpMethod string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       fmt.Sprintf("The method %s is not allowed for this endpoint.", httpMethod),
	}
}


func Created(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       body,
	}
}

func Ok(body string) events.APIGatewayProxyResponse  {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}
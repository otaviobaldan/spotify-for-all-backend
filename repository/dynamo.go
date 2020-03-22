package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	. "github.com/otaviobaldan/spotify-for-all-backend/models"
	uuid "github.com/satori/go.uuid"
)

const AWS_REGION = "us-east-1"
const TABLE_NAME = "spotify-for-all"
var db *dynamodb.DynamoDB

func init() {
	db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(AWS_REGION))
}

func CreateUser(user User) (*User, error) {
	// Generates a new random ID
	user.ID = uuid.NewV4().String()

	// Creates the item that's going to be inserted
	input := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.ID),
			},
			"name": {
				S: aws.String(user.Name),
			},
			"age": {
				S: aws.String(user.Age),
			},
		},
	}

	_, err := db.PutItem(input)
	return &user, err
}

func GetUsers() ([]User, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
	}
	result, err := db.Scan(input)
	if err != nil {
		return []User{}, err
	}
	if len(result.Items) == 0 {
		return []User{}, nil
	}

	var users []User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

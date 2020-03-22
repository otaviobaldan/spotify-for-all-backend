package service

import (
	"github.com/otaviobaldan/spotify-for-all-backend/models"
	"github.com/otaviobaldan/spotify-for-all-backend/repository"
)

func CreateUser(user models.User) (*models.User, error) {
	response, err := repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetUsers() ([]models.User, error) {
	response, err := repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return response, nil
}

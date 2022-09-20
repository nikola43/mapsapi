package services

import (
	"errors"

	database "github.com/nikola43/mapsapi/database"
	"github.com/nikola43/mapsapi/models"
	"github.com/nikola43/mapsapi/utils"
)

func LoginClient(email, password string) (*models.LoginClientResponse, error) {
	client := &models.User{}

	err := database.GormDB.
		Where("email = ?", email).
		Find(&client).Error
	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(client.Password, []byte(password)) {
		return nil, errors.New("user not found")
	}

	token, err := utils.GenerateClientToken(client.Email, client.ID)
	if err != nil {
		return nil, err
	}

	clientLoginResponse := &models.LoginUserResponse{
		Id:    client.ID,
		Email: client.Email,
		Name:  client.Name,
		Token: token,
	}

	return clientLoginResponse, err
}

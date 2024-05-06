package usecase

import "rest-api-example/internal/models"

type RegisterUser struct {
	Name     string
	Email    string
	Password string
}

func (d *RegisterUser) toUser() models.User {
	return models.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}

type SignIn struct {
	Email    string
	Password string
}

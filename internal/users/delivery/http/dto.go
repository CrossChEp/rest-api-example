package http

import "rest-api-example/internal/users/usecase"

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) toRegisterUser() usecase.RegisterUser {
	return usecase.RegisterUser{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *SignInRequest) toSignIn() usecase.SignIn {
	return usecase.SignIn{
		Email:    r.Email,
		Password: r.Password,
	}
}

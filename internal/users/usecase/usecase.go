package usecase

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"rest-api-example/pkg/utils"
)

type UserUC struct {
	cfg      *config.Config
	userRepo UserRepo
}

func NewUserUC(cfg *config.Config, userRepo UserRepo) *UserUC {
	return &UserUC{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u *UserUC) Register(ctx context.Context, regData RegisterUser) (models.UserID, error) {
	user := regData.toUser()

	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	user.Password = string(hashedPswd)
	return u.userRepo.Create(ctx, user)
}

func (u *UserUC) SignIn(ctx context.Context, signInData SignIn) error {
	user, err := u.userRepo.Get(ctx, models.UserFilter{
		Emails: []string{signInData.Email},
	})
	if err != nil {
		return err
	}

	if !utils.IsPasswordCorrect([]byte(signInData.Password), []byte(user.Password)) {
		return errors.New("wrong password")
	}

	return nil
}

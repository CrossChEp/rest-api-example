package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"rest-api-example/pkg/utils"
)

type UserUC struct {
	cfg           *config.Config
	userRepo      UserRepo
	userRedisRepo UserRedisRepo
}

func NewUserUC(cfg *config.Config, userRepo UserRepo, userRedisRepo UserRedisRepo) *UserUC {
	return &UserUC{
		cfg:           cfg,
		userRepo:      userRepo,
		userRedisRepo: userRedisRepo,
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

func (u *UserUC) SignIn(ctx context.Context, signInData SignIn) (string, error) {
	user, err := u.userRepo.Get(ctx, models.UserFilter{
		Emails: []string{signInData.Email},
	})
	if err != nil {
		return "", err
	}

	if !utils.IsPasswordCorrect([]byte(signInData.Password), []byte(user.Password)) {
		return "", errors.New("wrong password")
	}

	sessionKey := fmt.Sprintf("%d:%s", user.ID, uuid.NewString())

	if err := u.userRedisRepo.SetUserSession(ctx, sessionKey, models.Claims{
		UserID: user.ID,
		Email:  user.Email,
	}); err != nil {
		return "", err
	}

	return sessionKey, nil
}

func (u *UserUC) GetUser(ctx context.Context, userID models.UserID) (models.User, error) {
	return u.userRepo.Get(ctx, models.UserFilter{
		IDs: []models.UserID{userID},
	})
}

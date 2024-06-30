package usecase

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"rest-api-example/pkg/utils"
	"time"
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

	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = models.Claims{
		UserID:         user.ID,
		Email:          user.Email,
		ExpiresAt:      time.Now().Add(time.Minute * 30),
		StandardClaims: jwt.StandardClaims{},
	}

	tokenStr, err := token.SignedString(u.cfg.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (u *UserUC) RefreshToken(ctx context.Context, userID models.UserID) (string, error) {
	user, err := u.userRepo.Get(ctx, models.UserFilter{
		IDs: []models.UserID{userID},
	})
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = models.Claims{
		UserID:         user.ID,
		Email:          user.Email,
		ExpiresAt:      time.Now().Add(time.Minute * 30),
		StandardClaims: jwt.StandardClaims{},
	}

	tokenStr, err := token.SignedString(u.cfg.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (u *UserUC) GetUser(ctx context.Context, userID models.UserID) (models.User, error) {
	user, err := u.userRepo.Get(ctx, models.UserFilter{
		IDs: []models.UserID{userID},
	})
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

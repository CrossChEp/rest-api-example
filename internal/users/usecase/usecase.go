package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"rest-api-example/pkg/utils"
	"strconv"
	"time"
)

type UserUC struct {
	cfg       *config.Config
	userRepo  UserRepo
	userRedis UserRedisRepo
}

func NewUserUC(cfg *config.Config, userRepo UserRepo, userRedis UserRedisRepo) *UserUC {
	return &UserUC{
		cfg:       cfg,
		userRepo:  userRepo,
		userRedis: userRedis,
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

func (u *UserUC) SignIn(ctx context.Context, signInData SignIn) (models.Session, error) {
	user, err := u.userRepo.Get(ctx, models.UserFilter{
		Emails: []string{signInData.Email},
	})
	if err != nil {
		return models.Session{}, err
	}

	if !utils.IsPasswordCorrect([]byte(signInData.Password), []byte(user.Password)) {
		return models.Session{}, errors.New("wrong password")
	}

	session := models.Session{
		SessionKey: strconv.Itoa(int(user.ID)) + ":" + uuid.NewString() + ":user",
		TTL:        u.cfg.SessionSettings.SessionTTLSeconds,
	}

	if err := u.userRedis.PutSession(ctx, models.CacheUserSession{
		SessionKey: session.SessionKey,
		UserAgent:  "",
		Duration:   time.Duration(u.cfg.SessionSettings.SessionTTLSeconds),
		ID:         int(user.ID),
	}); err != nil {
		return models.Session{}, nil
	}

	return session, nil
}

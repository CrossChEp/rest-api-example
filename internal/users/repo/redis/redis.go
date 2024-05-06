package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"strings"
)

type UserRedisRepo struct {
	db  *redis.Client
	cfg *config.Config
}

func NewUserRedisRepo(db *redis.Client, cfg *config.Config) *UserRedisRepo {
	return &UserRedisRepo{db: db, cfg: cfg}
}

var SessionPrefix = "session"

func (r *UserRedisRepo) getDbKey(sessionKey, userAgent string) string {
	return strings.Join([]string{
		SessionPrefix,
		sessionKey,
		userAgent,
	}, ":")
}

// used for checking in middleware.
func (r *UserRedisRepo) getDbKeyByAuthHeaders(args models.AuthHeaders) string {
	return r.getDbKey(args.SessionKey, args.UserAgent)
}

// used for login.
func (r *UserRedisRepo) getDbKeyByCacheClientSession(args models.CacheUserSession) string {
	return r.getDbKey(args.SessionKey, args.UserAgent)
}

func (r *UserRedisRepo) GetSession(
	ctx context.Context,
	authHeaders models.AuthHeaders,
) (models.CacheUserSession, error) {
	result := models.CacheUserSession{}
	key := r.getDbKeyByAuthHeaders(authHeaders)

	valueString, err := r.db.Get(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return result, errors.New("UserRedisRepo.GetSession")
	} else if err != nil {
		return result, errors.New("UserRedisRepo.GetSession")
	}
	err = json.Unmarshal([]byte(valueString), &result)
	if err != nil {
		return result, errors.New("UserRedisRepo.GetSession")
	}
	return result, nil
}

func (r *UserRedisRepo) PutSession(ctx context.Context, session models.CacheUserSession) error {

	sessionBytes, err := json.Marshal(session)
	if err != nil {
		return err
	}

	key := r.getDbKeyByCacheClientSession(session)
	_, err = r.db.Set(ctx, key, sessionBytes, session.Duration).Result()
	if err != nil {
		return err
	}
	return nil
}

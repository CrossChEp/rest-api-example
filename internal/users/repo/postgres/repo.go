package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"rest-api-example/config"
	"rest-api-example/internal/models"
	"rest-api-example/pkg/sqlQueries"
	"time"
)

type UserRepo struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewUserRepo(cfg *config.Config, db *sqlx.DB) *UserRepo {
	return &UserRepo{
		cfg: cfg,
		db:  db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user models.User) (models.UserID, error) {
	query, args, err := sq.Insert(sqlQueries.UserTable).
		Columns(sqlQueries.InsertUserColumns...).
		Values(
			user.Name,
			user.Email,
			user.Password,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.UserIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	var userID models.UserID
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&userID)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (r *UserRepo) Get(ctx context.Context, filter models.UserFilter) (models.User, error) {
	users, err := r.GetMany(ctx, filter)
	if err != nil {
		return models.User{}, nil
	}

	if len(users) == 0 {
		return models.User{}, sql.ErrNoRows
	}
	return users[0], nil
}

func (r *UserRepo) GetMany(ctx context.Context, filter models.UserFilter) ([]models.User, error) {

	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries.GetUserColumns...).
		From(sqlQueries.UserTable).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var manyUsers manyUsers
	err = r.db.SelectContext(ctx, &manyUsers, query, args...)
	if err != nil {
		return nil, err
	}

	return manyUsers.toManyUsers(), nil
}

func (r *UserRepo) getConds(filter models.UserFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.UserIDColumnName: filter.IDs,
		})
	}
	if len(filter.Names) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.NameColumnName: filter.Names,
		})
	}
	if len(filter.Emails) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.EmailColumnName: filter.Emails,
		})
	}
	if len(filter.Passwords) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.PasswordColumnName: filter.Passwords,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries.DeletedAtColumnName: nil,
	})
	return sb
}

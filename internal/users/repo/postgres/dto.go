package postgres

import (
	"rest-api-example/internal/models"
	"time"
)

type manyUsers []user

func (d manyUsers) toManyUsers() []models.User {
	users := make([]models.User, 0, len(d))
	for _, u := range d {
		users = append(users, u.toUser())
	}
	return users
}

type user struct {
	ID        models.UserID `db:"id"`
	Name      string        `db:"name"`
	Email     string        `db:"email"`
	Password  string        `db:"password"`
	CreatedAt time.Time     `db:"create_at"`
	//DeletedAt time.Time     `db:"deleted_at"`
}

func (d *user) toUser() models.User {
	return models.User{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		Password:  d.Password,
		CreatedAt: d.CreatedAt,
		//DeletedAt: d.DeletedAt,
	}
}

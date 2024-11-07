package sql

import (
	_ "embed"
	"github.com/dafuqqqyunglean/music_library/pkg/models"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository interface {
	Create(user models.User) (int, error)
	Get(username, password string) (models.User, error)
}

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

//go:embed query/CreateUser.sql
var createUser string

func (r *AuthorizationPostgres) Create(user models.User) (int, error) {
	var id int

	row := r.db.QueryRow(createUser, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//go:embed query/GetUser.sql
var getUser string

func (r *AuthorizationPostgres) Get(username, password string) (models.User, error) {
	var user models.User

	err := r.db.Get(&user, getUser, username, password)

	return user, err
}

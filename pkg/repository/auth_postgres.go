package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"rest_api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostrgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user rest_api.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (rest_api.User, error) {
	var user rest_api.User
	query := fmt.Sprintf("SELECT * from %s where username=$1 and password_hash=$2", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}

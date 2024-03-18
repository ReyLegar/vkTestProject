package repository

import (
	"database/sql"
	"fmt"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var UserID int
	query := `INSERT INTO users (username, PasswordHash, role) values ($1, $2, $3) RETURNING UserID`

	err := r.db.QueryRow(query, user.Username, user.PasswordHash, user.Role).Scan(&UserID)
	if err != nil {
		return -1, err
	}

	return UserID, nil
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := `SELECT UserID, username, passwordhash, role FROM users WHERE username=$1 AND passwordhash=$2`
	err := r.db.QueryRow(query, username, password).Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil
}

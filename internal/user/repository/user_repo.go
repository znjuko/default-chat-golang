package repository

import (
	"database/sql"
	"main/internal/models"
)

type UserRepositoryRealisation struct {
	database *sql.DB
}

func NewUserRepositoryRealisation(db *sql.DB) UserRepositoryRealisation {
	return UserRepositoryRealisation{database: db}
}

func (User UserRepositoryRealisation) GetUserData(login string) (string, int, error) {

	var userId *int
	var pass *string

	row := User.database.QueryRow("SELECT u_id , password FROM users WHERE login = $1", login)

	err := row.Scan(&userId, &pass)

	if err != nil {
		return "", 0, err
	}

	return *pass, *userId, nil

}

func (User UserRepositoryRealisation) AddNewUser(newUserData models.Session) (int, error) {

	row := User.database.QueryRow("INSERT INTO users (login,password) VALUES($1,$2) RETURNING u_id", *newUserData.Login, *newUserData.Password)

	userId := 0

	err := row.Scan(&userId)

	return userId, err
}

package usecase

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"main/internal/cookie"
	"main/internal/models"
	"main/internal/user"
	"time"
)

type UserUseCaseRealisation struct {
	userDB    user.UserRepository
	sessionDB cookies.CookieRepository
}

func NewUserUseCaseRealisation(user user.UserRepository, sess cookies.CookieRepository) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:    user,
		sessionDB: sess,
	}
}

func (User UserUseCaseRealisation) Logout(cookie string) error {

	return User.sessionDB.DeleteCookie(cookie)

}

func (User UserUseCaseRealisation) Login(session models.Session) (string, error) {

	currentPass, userId, err := User.userDB.GetUserData(*session.Login)

	if err != nil {
		return "", err
	}

	if currentPass != *session.Password {
		return "", errors.New("wrong password")
	}

	cookieGen := uuid.NewV4()
	cookieValue := cookieGen.String()

	err = User.sessionDB.AddCookie(userId, cookieValue, 16*time.Hour)

	if err != nil {
		return "", err
	}

	return cookieValue, nil
}

func (User UserUseCaseRealisation) Register(newUser models.Session) (string, error) {

	userId, err := User.userDB.AddNewUser(newUser)

	if err != nil {
		return "", err
	}

	cookieGen := uuid.NewV4()
	cookieValue := cookieGen.String()

	err = User.sessionDB.AddCookie(userId, cookieValue, 16*time.Hour)

	if err != nil {
		return "", err
	}

	return cookieValue, nil

}

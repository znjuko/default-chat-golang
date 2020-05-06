package user

import "main/internal/models"

type UserUseCase interface {
	Logout(string) error
	Login(models.Session) (string, error)
	Register(models.Session) (string, error)
}
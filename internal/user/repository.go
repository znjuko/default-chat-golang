package user

import "main/internal/models"

type UserRepository interface {
	GetUserData(string) (string, int, error)
	AddNewUser(models.Session) (int, error)
}

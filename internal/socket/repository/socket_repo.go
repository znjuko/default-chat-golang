package repository

import (
	"database/sql"
	"main/internal/models"
)

type OnlineRepoRealistaion struct {
	database *sql.DB
}

func NewOnlineRepoRealosation(db *sql.DB) OnlineRepoRealistaion {
	return OnlineRepoRealistaion{database: db}
}

func (Socket OnlineRepoRealistaion) AddOnline(userId int) error {

	_, err := Socket.database.Exec("INSERT INTO online (u_id) VALUES ($1)", userId)

	return err

}

func (Socket OnlineRepoRealistaion) DiscardOnline(userId int) error {

	_, err := Socket.database.Exec("DELETE FROM online WHERE u_id = $1", userId)

	return err

}

func (Socket OnlineRepoRealistaion) GetOnline(userId int) ([]models.OnlineUsers,error) {
	row, err := Socket.database.Query("SELECT U.u_id , U.login FROM users U INNER JOIN online O ON (O.u_id=U.u_id) WHERE U.u_id != $1", userId)

	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	if err != nil {
		return nil, err
	}

	users := make([]models.OnlineUsers, 0)

	for row.Next() {
		user := new(models.OnlineUsers)

		err = row.Scan(&user.UserId, &user.Login)

		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return users, nil
}

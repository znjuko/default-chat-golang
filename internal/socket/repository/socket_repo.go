package repository

import "database/sql"

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

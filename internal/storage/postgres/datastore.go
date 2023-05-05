package postgres

import (
	"database/sql"
	"github.com/dobb2/zenTotem/internal/entity"
)

type UserStorer struct {
	db *sql.DB
}

func Create(db *sql.DB) *UserStorer {
	return &UserStorer{db: db}
}

func (u UserStorer) Create(user entity.User) (entity.User, error) {
	query := `
	INSERT INTO Client (name, age) VALUES($1, $2) RETURNING id;
	`
	var idUser entity.User
	err := u.db.QueryRow(query, user.Name, user.Age).Scan(&idUser.Id)

	if err != nil {
		return entity.User{}, err
	}

	return idUser, nil

}

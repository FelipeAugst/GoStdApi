package repository

import (
	"api/src/db"
	"api/src/models"
	"database/sql"
)

type UserRepository interface {
	Create(models.User) error
	Get() ([]models.User, error)
	Find(id uint64) (models.User, error)
	Update(models.User) error
	Delete(id uint64) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(models.User) error {
	panic("unimplemented")
}

func (u *userRepository) Delete(id uint64) error {
	panic("unimplemented")
}

func (u *userRepository) Find(id uint64) (models.User, error) {
	defer u.db.Close()
	result, err := u.db.Query(`select ID,Name,Email,Nickname from users
	where id=? `)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	for result.Next() {
		if err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Nickname, &user.Password); err != nil {
			return models.User{}, err
		}

	}

}

func (u *userRepository) Get() ([]models.User, error) {
	query := "Select ID,Name,Email,Nickname from users;"
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer u.db.Close()
	var users []models.User
	var user models.User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Nickname, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) Update(user models.User) error {
	defer u.db.Close()
	stmnt, err := u.db.Prepare(`Update Users
	set name=?,
	email=?,
	nickname=?
	where id=?`)
	if err != nil {
		return err
	}
	defer stmnt.Close()
	if _, err := stmnt.Exec(user.Name, user.Email, user.Nickname, user.ID); err != nil {
		return err
	}
	return nil

}

func NewUserRepository() (UserRepository, error) {
	db, err := db.Connect("")
	if err != nil {
		return &userRepository{}, err
	}
	return &userRepository{db: db}, nil
}

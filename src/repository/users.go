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

func (u *userRepository) Create(user models.User) error {
	defer u.db.Close()
	sttmnt, err := u.db.Prepare(
		"INSERT INTO USERS(NAME,EMAIL,NICKNAME,PASSWORD) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer sttmnt.Close()
	if _, err := sttmnt.Exec(user.Name, user.Email, user.Nickname, user.Password); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id uint64) error {
	_, err := u.db.Query(`DELETE FROM USERS
	WHERE ID=?`, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Find(id uint64) (models.User, error) {
	defer u.db.Close()
	result, err := u.db.Query(`SELECT ID,NAME,EMAIL,NICKNAME,PASSWORD FROM USERS
	WHERE ID=?`, id)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	for result.Next() {
		if err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Nickname, &user.Password); err != nil {
			return models.User{}, err
		}

	}
	return user, nil

}

func (u *userRepository) Get() ([]models.User, error) {
	query := "SELECT ID,NAME,EMAIL,NICKNAME,PASSWORD FROM USERS;"
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
	stmnt, err := u.db.Prepare(`UPDATE USERS
	SET NAME=?,
	EMAIL=?,
	NICKNAME=?
	WHERE ID=?`)
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
	db, err := db.Connect("felipe:felipe@/TEST")
	if err != nil {
		return &userRepository{}, err
	}
	return &userRepository{db: db}, nil
}

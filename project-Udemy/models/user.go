package models

import (
	"errors"

	"github.com/Wong-bui/Udemy-project/db"
	"github.com/Wong-bui/Udemy-project/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (user User) Save() error {
	database := db.GetInstance()
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := database.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}

func (user *User) ValidateCredentials() error {
	database := db.GetInstance()

	query := "SELECT id, password FROM users WHERE email = ?"
	row := database.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}

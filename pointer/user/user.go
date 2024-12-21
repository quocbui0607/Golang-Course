package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreateAt  time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

func PrintOutputUserDetails(u User) {
	fmt.Println(u.FirstName)
	fmt.Println(u.LastName)
	fmt.Println(u.BirthDate)
}

func GetUserData(text string) string {
	var input string
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}

// Receiver or receiver arguments
func (u User) PrintOutputUserDetails() {
	fmt.Println("Methods outputUserDetails")
	fmt.Print(u.FirstName, " ", u.LastName, "\n")
}

func (u *User) ClearUserDetails() {
	u.FirstName = ""
	u.LastName = ""
	u.BirthDate = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "----",
			CreateAt:  time.Now(),
		},
	}
}

// Creation / Construct
func NewUser(userFirstName string, userLastName string, userBirthDate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("First name, last name and birthdate is invalid")
	}

	return &User{
		FirstName: userFirstName,
		LastName:  userLastName,
		BirthDate: userBirthDate,
		CreateAt:  time.Now(),
	}, nil
}

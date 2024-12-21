package main

import (
	user "test.com/pointer/user"
)

// func main() {
// 	age := 32

// 	agePointer := &age

// 	fmt.Println("Age:", *agePointer) // dereference

// 	getAdultYears(agePointer)

// 	fmt.Println("Age:", age)
// }

// func getAdultYears(agePointer *int) {
// 	*agePointer = *agePointer - 18 // dereference
// }

func main() {
	userFirstName := user.GetUserData("Please enter your firstName: ")
	userLastName := user.GetUserData("Please enter your lastName: ")
	userBirthDate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	var appUser2 *user.User
	appUser, _ = user.NewUser(userFirstName, userLastName, userBirthDate)
	appUser2, _ = user.NewUser(userFirstName, userLastName, userBirthDate)
	adminUser := user.NewAdmin("test@test.com", "123456")

	adminUser.PrintOutputUserDetails()

	appUser.PrintOutputUserDetails()
	appUser.ClearUserDetails()
	user.PrintOutputUserDetails(*appUser)
	user.PrintOutputUserDetails(*appUser2)

}

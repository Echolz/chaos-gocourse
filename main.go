package main

import (
	"fmt"
	"github.com/Echolz/chaos-gocourse/homework3"
)

func main() {
	storage := homework3.NewStorage()

	err := storage.CreateUser(homework3.UserRequest{
		Username:  "echolz",
		Password:  "1234qwer",
		Email:     "echolz@abv.bg",
		FirstName: "echolz",
		LastName:  "echolz",
		UserRole:  "user",
	})
	logError(err)

	currentUser, err := storage.GetUser(1)
	logError(err)
	fmt.Println(currentUser.String())

	err = storage.UpdateUser(1, homework3.UserRequest{
		Username:  "newname",
		Password:  "1234qwer",
		Email:     "echolz@abv.bg",
		FirstName: "newname",
		LastName:  "newname",
		UserRole:  "user",
	})
	logError(err)

	currentUser, err = storage.GetUser(1)
	logError(err)
	fmt.Println(currentUser.String())

	err = storage.DeleteUser(1)
	logError(err)

	currentUser, err = storage.GetUser(1)
	logError(err)
	fmt.Println(currentUser.String())
}

func logError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import "github.com/Echolz/chaos-gocourse/homework3"

const localHost string = "http://localhost:8080"

func main() {
	homework3.SortByRequest(localHost, homework3.ByFirstNameString)
	//homework3.CreateUserRequest(localHost+"/users", homework3.UserRequest{
	//	Username:  "",
	//	Password:  "",
	//	Email:     "",
	//	FirstName: "ccc",
	//	LastName:  "",
	//	UserRole:  "admin",
	//})
	//
	//homework3.SortByRequest(localHost, homework3.ByFirstNameString)
	//homework3.SortByRequest(localHost, homework3.ByRoleString)
	//homework3.GetUser(localHost, "1")
	//homework3.DeleteUser(localHost, "1")
	//homework3.GetUser(localHost, "1")
	homework3.UpdateUser(localHost, "0", homework3.UserRequest{
		Username:  "",
		Password:  "",
		Email:     "",
		FirstName: "ccc",
		LastName:  "",
		UserRole:  "user",
	})
	homework3.GetUser(localHost, "0")
}

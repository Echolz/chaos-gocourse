package homework3

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SortByRequest(url, field string) {
	url = fmt.Sprintf("%s/users?sortBy=%s", url, field)
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	var sortedUsers []User

	err = json.NewDecoder(resp.Body).Decode(&sortedUsers)

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range sortedUsers {
		fmt.Println(user.String())
	}
}

func CreateUserRequest(url string, request UserRequest) {

}

func GetUser(url, id string) {

}

func UpdateUser(url, id string, request UserRequest) {

}

func DeleteUser(url, id string) {

}

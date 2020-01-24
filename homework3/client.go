package homework3

import (
	"bytes"
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

	defer resp.Body.Close()

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
	reqBody, err := json.Marshal(request)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
}

func GetUser(url, id string) {
	url = fmt.Sprintf("%s/users/%s", url, id)
	fmt.Println("GET: ", url)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("User %s not found\n", id)
		return
	}

	var user User

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.String())
}

func UpdateUser(url, id string, request UserRequest) {
	url = fmt.Sprintf("%s/users/%s", url, id)
	client := &http.Client{}

	reqBody, err := json.Marshal(request)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	fmt.Println("response Status : ", resp.Status)
}

func DeleteUser(url, id string) {
	url = fmt.Sprintf("%s/users/%s", url, id)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	fmt.Println("response Status : ", resp.Status)
}

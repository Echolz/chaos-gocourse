package homework3

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const badID int = -1

type Handler interface {
	SortBy(w http.ResponseWriter, req *http.Request)
	CreateUser(w http.ResponseWriter, req *http.Request)
	GetUser(w http.ResponseWriter, req *http.Request)
	UpdateUser(w http.ResponseWriter, req *http.Request)
	DeleteUser(w http.ResponseWriter, req *http.Request)
}

type concreteHandler struct {
	storage storage
}

func (c concreteHandler) SortBy(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	field := req.URL.Query().Get("sortBy")

	users := c.storage.sortBy(field)

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (c concreteHandler) CreateUser(w http.ResponseWriter, req *http.Request) {

}

func (c concreteHandler) GetUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v, ok := extractId(req)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := c.storage.getUser(v)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (c concreteHandler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (c concreteHandler) DeleteUser(w http.ResponseWriter, req *http.Request) {
	id, ok := extractId(req)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := c.storage.deleteUser(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func extractId(req *http.Request) (int, bool) {
	params := mux.Vars(req)

	str, ok := params["id"]

	if !ok {
		return badID, false
	}

	v, err := strconv.Atoi(str)

	if err != nil {
		return badID, false
	}

	return v, true
}

func NewHandler() Handler {
	h := &concreteHandler{}
	h.storage = newStorage()

	h.storage.createUser(UserRequest{
		Username:  "",
		Password:  "",
		Email:     "",
		FirstName: "aa",
		LastName:  "",
		UserRole:  "user",
	})

	return h
}

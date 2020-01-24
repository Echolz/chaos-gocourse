package main

import (
	"github.com/Echolz/chaos-gocourse/homework3"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	h := homework3.NewHandler()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", h.SortBy).Methods(http.MethodGet)
	router.HandleFunc("/users", h.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

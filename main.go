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
	router.HandleFunc("/posts", h.CreateUser).Methods("POST")
	router.HandleFunc("/posts/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/posts/{id}", h.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

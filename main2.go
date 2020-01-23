package main

import "github.com/Echolz/chaos-gocourse/homework3"

const localHost string = "http://localhost:8080"

func main() {
	homework3.SortByRequest(localHost, homework3.ByFirstNameString)
}

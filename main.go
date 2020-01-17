package main

import (
	"github.com/Echolz/chaos-gocourse/homework1"
	"github.com/Echolz/chaos-gocourse/lab2"
	"log"
	"os/user"
)

func main() {

	//runFirstHomework()

	//runSecondLab()

	//runThirdLab()
}

func runFirstHomework() {
	homework1.Run()
}

func runSecondLab() {
	u, err := user.Current()

	if err != nil {
		log.Fatal(err.Error())
	}

	if u == nil {
		log.Fatal("Couldn't get user")
	}

	lab2.Run(u.HomeDir, "sucks.jpg")
}

package main

import (
	"fmt"
	"time"
)

type Guest struct {
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}

func newGuest(firstName string, lastName string, dateOfBirth time.Time) *Guest {
	return &Guest{
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
	}
}

func main() {
	scooter := newGuest("Ryuichiro", "Suzuki", time.Date(2000, 3, 19, 9, 0, 0, 0, time.Local))
	fmt.Println("hello")
	fmt.Println(scooter)
	fmt.Printf("%p", scooter)
}

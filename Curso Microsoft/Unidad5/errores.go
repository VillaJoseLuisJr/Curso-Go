package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Adress    string
}

func main() {
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
	} else {
		fmt.Print(employee)
	}
}

var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

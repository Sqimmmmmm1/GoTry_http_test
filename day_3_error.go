package main

import (
	"errors"
)

func divede(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot dive by zero")
	}
	return a / b, nil
}

func getUserByID(id string) (string, error) {
	if id == "" {
		return "", errors.New("id is required")
	}
	if id == "1" {
		return "Amy", nil
	}
	if id == "2" {
		return "Bob", nil
	}

	return "", errors.New("user not found")

}

func checkAge(age int) error {
	if age < 0 {
		return errors.New("age is negative")
	}
	return nil
}

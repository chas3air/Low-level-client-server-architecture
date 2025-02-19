package models

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
	Role     string
	Nick     string
}

func NewUser() *User {
	id := uuid.New()
	var email, password, role, nick string
	fmt.Println("Enter email")
	fmt.Scanf("%s", &email)

	fmt.Println("Enter password")
	fmt.Scanf("%s", &password)

	fmt.Println("Enter role")
	fmt.Scanf("%s", &role)

	fmt.Println("Enter nick")
	fmt.Scanf("%s", &nick)

	return &User{
		Id:       id,
		Email:    email,
		Password: password,
		Role:     role,
		Nick:     nick,
	}
}

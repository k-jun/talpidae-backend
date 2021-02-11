package user

import "github.com/google/uuid"

type User struct {
	id   uuid.UUID
	name string
}

func New(id uuid.UUID, name string) *User {
	return &User{id, name}
}

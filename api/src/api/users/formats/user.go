package formats

import (
	"github.com/google/uuid"
)

/*
User structure
*/
type User struct {
	Name          string
	Email         string
	Admin         bool
	Password      Password
	Notifications bool
	ID            uuid.UUID
	LastLogin     uint64
	CreatedAt     uint64
}

/*
Password structure
*/
type Password struct {
	Encryption string
	Vector     string
	Password   string
}

package formats

import (
	"time"

	"github.com/google/uuid"
)

/*
User structure
*/
type User struct {
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Admin         bool      `json:"admin"`
	Password      string    `json:"password"`
	Notifications bool      `json:"notifications"`
	ID            uuid.UUID `json:"id"`
	LastLogin     time.Time `json:"lastLogin"`
	CreatedAt     time.Time `json:"createdAt"`
}

/*
Password structure
*/
type Password struct {
	Password string `json:"password"`
}

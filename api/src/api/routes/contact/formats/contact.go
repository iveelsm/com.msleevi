package formats

import (
	"time"
)

type Contact struct {
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

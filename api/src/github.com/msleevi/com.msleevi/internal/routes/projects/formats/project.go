package projects

import (
	"time"
)

/*
Project does
*/
type Project struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	URL         string    `json:"url"`
	Image       string    `json:"image"`
	Synopsis    string    `json:"synopsis"`
	Description string    `json:"description"`
}

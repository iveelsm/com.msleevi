package projects

import (
	"time"
)

/*
Project does
*/
type Project struct {
	name        string
	date        time.Time
	url         string
	image       string
	synopsis    string
	description string
}

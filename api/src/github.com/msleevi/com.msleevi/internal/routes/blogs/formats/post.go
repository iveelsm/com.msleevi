package formats

import (
	"time"

	"github.com/google/uuid"
)

/*
Blog structure
Blog test
*/
type Blog struct {
	Date     time.Time `json:"date"`
	Post     Post      `json:"post"`
	Stats    []Stat    `json:"stats"`
	Author   uuid.UUID `json:"author"`
	Comments []Comment `json:"comments"`
	Views    int       `json:"views"`
}

/*
Stat structure
*/
type Stat struct {
	Votes []Vote `json:"votes"`
}

/*
Comment structure
*/
type Comment struct {
	Author  uuid.UUID    `json:"author"`
	Date    time.Time    `json:"date"`
	Comment string       `json:"comment"`
	Replies []Comment    `json:"replies"`
	Stats   CommentStats `json:"stats"`
}

/*
Post structure
*/
type Post struct {
	Content string `json:"content"`
}

/*
CommentStats structure
*/
type CommentStats struct {
	Votes []Vote `json:"votes"`
}

/*
Vote structure
*/
type Vote struct {
	Vote  int       `json:"vote"`
	Voter uuid.UUID `json:"voter"`
	Date  time.Time `json:"date"`
}

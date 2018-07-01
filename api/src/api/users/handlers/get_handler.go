package handlers

import (
	"api/users/formats"

	"github.com/google/uuid"
)

func HandleUsers(params map[string]string) (users []formats.User, err error) {
	user := formats.User{
		"michael.sleevi@gmail.com",
		"Michael Sleevi",
		true,
		formats.Password{"test", "fdgdgs", "werqwrqwer"},
		true,
		uuid.New(),
		0,
		0,
	}
	return []formats.User{user}, nil
}

func HandleUser() {

}

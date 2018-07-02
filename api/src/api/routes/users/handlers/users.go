package handlers

import (
	"api/routes"
	"api/routes/users/formats"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := handleUsers(params)

	if err != nil {
		routes.HandleError("Unable to GET /users", 500, w)
		return
	}

	bytes, err := json.Marshal(result)

	if err != nil {
		routes.HandleError("Unable to parse data into return body", 504, w)
		return
	}

	w.Write(bytes)
	return
}

func handleUsers(params map[string]string) (users []formats.User, err error) {
	user := getUserStub()
	return []formats.User{user}, nil
}

func getUserStub() (result formats.User) {
	result = formats.User{
		Email: "michael.sleevi@gmail.com",
		Name:  "Michael Sleevi",
		Admin: true,
		Password: formats.Password{
			Encryption: "test",
			Vector:     "fdgdgs",
			Password:   "werqwrqwer",
		},
		Notifications: true,
		ID:            uuid.New(),
		LastLogin:     0,
		CreatedAt:     0,
	}
	return
}

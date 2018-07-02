package handlers

import (
	"api/routes/users/formats"
	"api/utils"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetUsers does
func GetUsers(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	result, err := handleUsers(params)

	if err != nil {
		return nil, 500, errors.New("Unable to GET /users")
	}

	bytes, err := json.Marshal(result)

	if err != nil {
		return nil, 504, errors.New("Unable to parse data into return body")
	}

	return bytes, 200, nil
}

// GetUser does
func GetUser(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	if _, exist := params["id"]; !exist {
		return nil, 416, errors.New("Route requires ID specifications")
	}

	result, err := handleUsers(params)

	if err != nil {
		return nil, 500, errors.New("Unable to GET /users")
	}

	bytes, err := json.Marshal(result)

	if err != nil {
		return nil, 504, errors.New("Unable to parse data into return body")
	}

	return bytes, 200, nil
}

func handleUsers(params map[string]string) (users []formats.User, err error) {
	user := getUserStub()
	return []formats.User{user}, nil
}

func getUserStub() (result formats.User) {
	password, err := utils.GeneratePassword("test")
	if err != nil {
		panic("Can't generate user stub?")
	}

	result = formats.User{
		Email:         "michael.sleevi@gmail.com",
		Name:          "Michael Sleevi",
		Admin:         true,
		Password:      password,
		Notifications: true,
		ID:            uuid.New(),
		LastLogin:     time.Now(),
		CreatedAt:     time.Now(),
	}
	return
}

package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/msleevi/com.msleevi/internal/database"
	"github.com/msleevi/com.msleevi/internal/routes/users/formats"
	"github.com/msleevi/com.msleevi/internal/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// GetUsers does
func GetUsers(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	result, err := handleUsers(params)

	if err != nil {
		return nil, 500, errors.New("Unable to GET /users: " + err.Error())
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
	return getUsers()
}

func sortOnName(users []formats.User) []formats.User {
	sort.Slice(users[:], func(i, j int) bool {
		return users[i].Name < users[j].Name
	})
	return users
}

func getUsers() (users []formats.User, err error) {
	response, err := database.GetAll("users")
	if err != nil {
		return []formats.User{}, errors.New("Unable to contact database")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []formats.User{}, errors.New("Unable to parse body")
	}

	log.Debug("Body: " + string(body))

	err = json.Unmarshal([]byte(body), users)
	if err != nil {
		return []formats.User{}, errors.New("Unable to parse into Users")
	}

	return users, err
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

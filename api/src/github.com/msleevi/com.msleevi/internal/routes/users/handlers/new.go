package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/msleevi/com.msleevi/internal/routes/users/formats"

	"github.com/gorilla/mux"
)

// NewUser does
func NewUser(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, 504, errors.New("Unable to parse into User object")
	}

	user := formats.User{}
	err = json.Unmarshal([]byte(body), user)

	if err != nil {
		return nil, 504, errors.New("Unable to construct User object")
	}

	result, err := handleNewUser(params, user)

	if err != nil {
		return nil, 500, errors.New("Unable to create new User")
	}

	return []byte(result), 204, nil
}

func handleNewUser(params map[string]string, user formats.User) (result string, err error) {
	result = "Successful post"
	err = nil
	return
}

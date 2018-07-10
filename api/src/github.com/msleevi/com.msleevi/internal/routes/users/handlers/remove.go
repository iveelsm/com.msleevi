package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// RemoveUser does
func RemoveUser(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	result, err := removeUser(params)

	if err != nil {
		return nil, 504, errors.New("Unable to delete user")
	}

	return []byte(result), 204, nil
}

func removeUser(params map[string]string) (result string, err error) {
	return "deletion succesful", nil
}

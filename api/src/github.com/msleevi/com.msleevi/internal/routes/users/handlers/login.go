package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/msleevi/com.msleevi/internal/routes/users/formats"
)

/*
LoginUser does
*/
func LoginUser(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, 504, errors.New("Unable to parse into Login object")
	}

	login := formats.Login{}
	err = json.Unmarshal([]byte(body), login)

	if err != nil {
		return nil, 504, errors.New("Unable to construt Login object")
	}

	result, err := handleLogin(login)

	if err != nil {
		return nil, 500, errors.New("Unable to login")
	}

	return []byte(result), 200, nil
}

func handleLogin(login formats.Login) (result string, err error) {
	result = "login successful"
	err = nil
	return
}

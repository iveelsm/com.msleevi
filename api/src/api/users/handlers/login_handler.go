package handlers

import (
	"api/users/formats"
	"api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func LoginUsersHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleError("Unable to parse into Login object", 504, w)
		return
	}

	login := formats.Login{}
	err = json.Unmarshal([]byte(body), login)

	if err != nil {
		utils.HandleError("Unable to construct Login object", 504, w)
		return
	}

	result, err := handleLogin(login)

	if err != nil {
		utils.HandleError("Unable to login", 500, w)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(result))
}

func handleLogin(login formats.Login) (result string, err error) {
	result = "login successful"
	err = nil
	return
}

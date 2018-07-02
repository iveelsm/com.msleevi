package handlers

import (
	"api/routes/users/formats"
	"api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func NewUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.HandleError("Unable to parse into User object", 504, w)
		return
	}

	user := formats.User{}
	err = json.Unmarshal([]byte(body), user)

	if err != nil {
		utils.HandleError("Unable to construct User object", 504, w)
		return
	}

	result, err := handleNewUser(params, user)

	if err != nil {
		utils.HandleError("Unable to create new User", 500, w)
		return
	}

	w.WriteHeader(204)
	w.Write([]byte(result))
}

func handleNewUser(params map[string]string, user formats.User) (result string, err error) {
	result = "Successful post"
	err = nil
	return
}

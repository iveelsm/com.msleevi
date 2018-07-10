package database

import (
	"net/http"

	"github.com/msleevi/com.msleevi/internal/constants"

	log "github.com/sirupsen/logrus"
)

func GetValue() {

}

func GetAll(database string) (*http.Response, error) {
	//stopwatch := time.Now()
	result, err := http.Get(constants.DATABASE_URI + "/" + database + "/_all_docs")
	if err != nil {
		log.Error("Unable to interact with database at: " + constants.DATABASE_URI)
		return result, err
	}

	return result, err
}

func PostValue() {

}

func DeleteValue() {

}

package userHandler

import (
	"auth/pkg/user"
	"encoding/json"
	"log"
	"net/http"
)

/**
Return all users from database
*/
func AllUsers(rw http.ResponseWriter, req *http.Request) {
	users := user.NewRepository(make([]user.User, 0)).All()

	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(users)

	if err != nil {
		log.Fatal(err.Error())
	}
}

package auth

import (
	database "auth/pkg/db"
	"database/sql"
	"log"
	"net/http"
	"strings"
)

var allowedHeaders = map[string]string{
	"Content-Type": "application/json",
	"Accept":       "application/json",
}

/**
	Auth middleware
*/
func AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {

		if !IsAuthorized(request) {
			http.Error(rw, "Unauthorized", http.StatusForbidden)
			return
		}

		if !ValidateHeaders(request) {
			http.Error(rw, "Bad Request", http.StatusBadRequest)
			return
		}

		handler(rw, request)
	}
}

/**
	Check if request is authorized
*/
func IsAuthorized(request *http.Request) bool {
	authHeaderValue, exist := request.Header["Authorization"]

	if !exist && len(authHeaderValue) != 1 {
		return false
	}

	splitBearer := strings.Split(authHeaderValue[0], "Bearer ")

	if len(splitBearer) != 2 {
		return false
	}

	var token string = splitBearer[1]

	var queryResult string
	err := database.Con.QueryRow("select token from users where token = ?", token).Scan(&queryResult)

	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	return true
}

/**
	Validate HTTP headers
*/
func ValidateHeaders(request *http.Request) bool {
	for key, value := range allowedHeaders {
		header, exists := request.Header[key]

		if !exists || len(header) != 1 || header[0] != value {
			return false
		}
	}

	return true
}

package main

import (
	"auth/pkg/auth"
	database "auth/pkg/db"
	userHandler "auth/pkg/user/handler"
	"log"
	"net/http"
)

func main() {
	database.Open()

	http.HandleFunc("/users/", auth.AuthMiddleware(userHandler.AllUsers))

	log.Fatal("HTTP server error: ", http.ListenAndServe(":8080", nil))
}

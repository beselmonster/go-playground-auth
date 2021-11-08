package main

import (
	"auth/pkg/auth"
	database "auth/pkg/db"
	userHandler "auth/pkg/user/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	port := "8080"

	database.Open()

	log.Printf("AUTH SERVER IS STARTING ON PORT %s ....", port)

	http.HandleFunc("/users/", auth.AuthMiddleware(userHandler.AllUsers))
	//http.HandleFunc("/users/", userHandler.AllUsers)

	log.Fatal("HTTP server error: ", http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

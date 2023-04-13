package user

import (
	"github.com/gorilla/mux"
	"github.com/isaqib23/golang-coffee-shop-microservices/database"
	"net/http"
)

func SetupRoutes(dbClient *database.DBClient, router *mux.Router) {
	userHandler := NewUserHandler(dbClient)

	// Get Routes
	getRoutes := router.Methods(http.MethodGet).Subrouter()
	getRoutes.HandleFunc("/", userHandler.FetchByName())
}

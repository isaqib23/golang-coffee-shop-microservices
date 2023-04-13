package user

import (
	"github.com/isaqib23/golang-coffee-shop-microservices/database"
)

func NewUserHandler(db *database.DBClient) *handler {
	userRepository := ProvideRepository(db)
	userService := ProvideService(userRepository)
	userHandler := ProvideHandler(userService)
	return userHandler
}

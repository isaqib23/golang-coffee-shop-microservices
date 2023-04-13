package config

import (
	"github.com/gorilla/mux"
	"github.com/isaqib23/golang-coffee-shop-microservices/database"
	"log"
	"net/http"
)

func NewIoC() IoC {
	return &IoCProfile{}
}

type IoC interface {
	Port() string
	MuxRouter() *mux.Router
	DbClient() *database.DBClient
	Logger() *log.Logger
	BindAddress() string
	HttpServer(router *mux.Router) *http.Server
}

package config

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewIoC() IoC {
	return &IoCProfile{}
}

type IoC interface {
	Port() string
	MuxRouter() *mux.Router
	HttpServer() http.Server
}

package config

import (
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"log"
	"net/http"
	"os"
	"time"
)

type IoCProfile struct{}

func (i *IoCProfile) Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = i.defaultPort()
	}
	return port
}

func (i *IoCProfile) defaultPort() string {
	return "8080"
}

func (i *IoCProfile) HttpServer() http.Server {
	return http.Server{
		Addr:         *i.bindAddress(),  // configure the bind address
		Handler:      i.MuxRouter(),     // set the default handler
		ErrorLog:     i.logger(),        // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
}

func (i *IoCProfile) bindAddress() *string {
	return env.String("BIND_ADDRESS", false, "localhost:8080", "BIND ADDRESS OF SERVER")
}

func (i *IoCProfile) logger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func (i *IoCProfile) MuxRouter() *mux.Router {
	return mux.NewRouter()
}

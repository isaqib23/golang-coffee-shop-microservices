package config

import (
	"github.com/gorilla/mux"
	"github.com/isaqib23/golang-coffee-shop-microservices/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type IoCProfile struct{}

func (i *IoCProfile) Port() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = i.defaultPort()
	}
	return port
}

func (i *IoCProfile) defaultPort() string {
	return "8080"
}

func (i *IoCProfile) BindAddress() string {
	return i.serverHost() + ":" + i.Port()
}

func (i *IoCProfile) Logger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func (i *IoCProfile) MuxRouter() *mux.Router {
	return mux.NewRouter()
}

func (i *IoCProfile) serverHost() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv("SERVER_HOST")
}

func (i *IoCProfile) DbClient() *database.DBClient {
	return database.NewDatabaseConnection()
}

func (i *IoCProfile) HttpServer(handler *mux.Router) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         i.BindAddress(),
		ErrorLog:     i.Logger(),        // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
}

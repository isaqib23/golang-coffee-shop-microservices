package main

import (
	"context"
	"github.com/isaqib23/golang-coffee-shop-microservices/config"
	"github.com/isaqib23/golang-coffee-shop-microservices/modules/user"
	"github.com/nicholasjackson/env"
	"log"
	"os"
	"os/signal"
	"time"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":8081", "BIND ADDRESS OF SERVER")

func main() {
	ioc := config.NewIoC()
	l := ioc.Logger()

	sm := ioc.MuxRouter()

	// User handler routes
	user.SetupRoutes(ioc.DbClient(), sm)

	// create a new server
	server := ioc.HttpServer(sm)

	// start the server
	go func() {
		l.Println("Starting server on port " + ioc.Port())

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}

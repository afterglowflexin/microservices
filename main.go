package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/afterglowflexin/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// creating handlers
	ph := handlers.NewProducts(l)

	// creating server multiplexer and registering handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// creating new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write the response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// starting the server
	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown.", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second) // gives 30 seconds
	s.Shutdown(tc)                                                     //does not accept new connections, waits until active connections finish
}

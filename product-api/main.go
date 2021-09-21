package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"working/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new logger. Is responsible for logging events.
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// Instantiate the handlers
	ph := handlers.NewProducts(l)

	// Creates a new Router instance. Router registers
	// routes to be matched and dispatches a handler.
	sm := mux.NewRouter()

	// Register a Get route and create its subrouter.
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	// Register products subroute and its associated handler.
	getRouter.HandleFunc("/products", ph.GetProducts)

	// Register a PUT route and create its subrouter.
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	// Register a products/:id subroute and its associated handler.
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	// Register a POST route and create its subrouter.
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	// Register a products/:id subroute and its associated handler.
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// Instantiate server at port 9090 and invoke router instance.
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Start the server.
	go func() {
		// Listen on port and direct inbound requests.
		err := s.ListenAndServe()
		l.Println("Starting server on port 9090")
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block until a signal is received
	sig := <-sigChan
	l.Println("Recieved Terminate, gracefuls shutdown", sig)

	// graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vitorcarra/product-api/handlers"
)

// var pl = fmt.Println

func main() {
	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	l.Println("Initializing server...")

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// starting the server using go routine
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// configuring gracefully shutdown to avoid issues with databases and ongoing processes
	sigChan := make(chan os.Signal)
	// we are going to broadcast to the channel when one of the signals happens
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// timeout context
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

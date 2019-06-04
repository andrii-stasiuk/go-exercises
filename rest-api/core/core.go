/*Package core*/
package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

// NewServer - creates and returns Server
func NewServer(listenAddr *string, router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

// StartServer - starts listen on Server
func StartServer(listenAddr *string, server *http.Server) {
	fmt.Println("Server is ready to handle requests at", *listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *listenAddr, err)
	}
}

// ShutdownServer - gracefull shutdown of the server
func ShutdownServer(server *http.Server, quit <-chan os.Signal, done chan<- struct{}) {
	// Waiting for SIGINT (pkill -2)
	<-quit
	// We received an interrupt signal, shut down
	fmt.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

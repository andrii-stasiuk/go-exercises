/*Package core to support the main application*/
package core

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// DatabaseConnect func creates and returnes new db (reserved for future purposes - to use with connection parameters)
func DatabaseConnect(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DatabaseVersion function that checks the operation of the database server and returns it version number
func DatabaseVersion(db *sql.DB) (string, error) {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive. Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		return "", err
	}

	var result string
	// Run query and scan for result with version number of PostgreSQL
	err = db.QueryRowContext(ctx, "SELECT version();").Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

// NewServer function that creates and returns http.Server
func NewServer(listenAddr *string, router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

// StartServer function that starts listen on server:port
func StartServer(listenAddr *string, server *http.Server) {
	fmt.Println("Server is ready to handle requests at", *listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *listenAddr, err)
	}
}

// ShutdownServer function that gracefully shutdown http.Server
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

// HashPassword - hash passwords using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - check password hash using bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

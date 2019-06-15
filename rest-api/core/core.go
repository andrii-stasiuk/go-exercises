/*Package core to support the main application*/
package core

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// DatabaseConnect func creates and returnes new db (reserved for future purposes - to use with connection parameters)
func DatabaseConnect(driverName, dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, nil
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
	log.Println("Server is ready to handle requests at", *listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *listenAddr, err)
	}
}

// ShutdownServer function that gracefully shutdown http.Server
func ShutdownServer(server *http.Server, quit <-chan os.Signal, done chan<- struct{}) {
	// Waiting for SIGINT (pkill -2)
	<-quit
	// We received an interrupt signal, shut down
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

// CheckInt function for basic verification of numbers, can be extended in the future
func CheckInt(id string) bool {
	converted, err := strconv.ParseUint(id, 10, 64)
	if err == nil && converted > 0 {
		return true
	}
	return false
}

// CheckStr function for basic string checking, can be extended in the future
func CheckStr(str string) bool {
	if len(strings.TrimSpace(str)) != 0 {
		return true
	}
	return false
}

// HashPassword - hash passwords using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash - check password hash using bcrypt
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/andrii-stasiuk/go-exercises/rest-api/handler"
	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/router"

	//	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func newServer(listenAddr *string, router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:         *listenAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

func gracefullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- struct{}) {
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

func main() {
	var dbURLPtr = flag.String("db", "root:@tcp(127.0.0.1:3306)/testdb", "Specify the URL to the database")
	var addrPtr = flag.String("addr", "127.0.0.1:8000", "Server IPv4 address")
	flag.Parse()

	fmt.Println("Server is starting...")

	//dataBase, err := sql.Open("postgres", "testuser:testpass@tcp(localhost:5555)/testdb?sslmode=disable")
	dataBase, err := model.DatabaseConnect("mysql", *dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	dataBase.SetMaxIdleConns(100)
	defer dataBase.Close()

	sql := model.Model{Db: dataBase}
	sqlVersion, err := sql.GetVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SQL Server version: %s\n", sqlVersion)

	srv := newServer(addrPtr, router.NewRouter(router.AllRoutes(handler.Handlers{SQL: sql})))

	done := make(chan struct{}, 1)
	// Setting up signal capturing
	quit := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quit, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(quit, syscall.SIGTERM)

	go gracefullShutdown(srv, quit, done)

	fmt.Println("Server is ready to handle requests at", *addrPtr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", *addrPtr, err)
	}

	<-done
	fmt.Println("Server gracefully stopped")
}

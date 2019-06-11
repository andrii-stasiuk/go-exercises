package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/router"
	"github.com/andrii-stasiuk/go-exercises/rest-api/todohandlers"
	"github.com/andrii-stasiuk/go-exercises/rest-api/todomodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/userhandlers"
	"github.com/andrii-stasiuk/go-exercises/rest-api/usermodel"
	_ "github.com/lib/pq"
)

func main() {
	// var dbURLPtr = flag.String("db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "Specify the URL to the database") // work DB
	var dbURLPtr = flag.String("db", "postgres://postgres:@localhost:5432/postgres?sslmode=disable", "Specify the URL to the database") // home DB
	var addrPtr = flag.String("addr", "127.0.0.1:8000", "Server IPv4 address")
	flag.Parse()

	fmt.Println("Server is starting...")

	dataBase, err := core.DatabaseConnect("postgres", *dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	dataBase.SetMaxIdleConns(100)
	defer dataBase.Close()

	todoModel := todomodel.New(dataBase)
	userModel := usermodel.New(dataBase)
	sqlVersion, err := core.DatabaseVersion(dataBase)
	// Checks the operation of the database server and returns it version number
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SQL Server version: %s\n", sqlVersion)

	srv := core.NewServer(addrPtr, router.NewRouter(router.AllRoutes(todohandlers.New(&todoModel), userhandlers.New(&userModel))))

	done := make(chan struct{}, 1)
	// Setting up signal capturing
	quit := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quit, os.Interrupt)

	go core.ShutdownServer(srv, quit, done)

	core.StartServer(addrPtr, srv)

	<-done
	fmt.Println("Server gracefully stopped")
}

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers/todo"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers/user"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/todomodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/router"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var dbURLPtr, addrPtr string

func init() {
	flag.StringVar(&dbURLPtr, "db", "postgres://postgres:@localhost:5432/postgres?sslmode=disable", "Specify the URL to the database")
	flag.StringVar(&addrPtr, "addr", "127.0.0.1:8000", "Server IPv4 address")
	flag.Parse()
}

func main() {
	log.Println("Server is starting...")

	db, err := gorm.Open("postgres", dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()

	log.Println("Successfully connected to Database")
	db.Debug().AutoMigrate(&todomodel.Todo{})

	dataBase, err := core.DatabaseConnect("postgres", dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	dataBase.SetMaxIdleConns(100)
	defer dataBase.Close()

	todoModel := todomodel.New(db)
	userModel := usermodel.New(dataBase)

	sqlVersion, err := core.DatabaseVersion(dataBase)
	// Checks the operation of the database server and returns it version number
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("SQL Server version: %s\n", sqlVersion)

	newRouer := router.NewRouter(
		router.TodoRoutes(todo.New(&todoModel)),
		router.UserRoutes(user.New(&userModel)))
	srv := core.NewServer(&addrPtr, newRouer)

	done := make(chan struct{}, 1)
	// Setting up signal capturing
	quit := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quit, os.Interrupt)

	go core.ShutdownServer(srv, quit, done)

	core.StartServer(&addrPtr, srv)

	<-done
	log.Println("Server gracefully stopped")
}

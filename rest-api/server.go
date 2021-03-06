package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/andrii-stasiuk/go-exercises/rest-api/common"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers/todo"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers/user"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/todomodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/router"
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

	dataBase, err := common.DBConnectSQLX("postgres", dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer dataBase.Close()

	log.Println("Successfully connected to Database")

	todoModel := todomodel.NewTodo(dataBase)
	userModel := usermodel.NewUser(dataBase)

	newRouer := router.NewRouter(
		router.TodoRoutes(todo.New(&todoModel)),
		router.UserRoutes(user.New(&userModel)))
	srv := common.NewServer(&addrPtr, newRouer)

	done := make(chan struct{}, 1)
	// Setting up signal capturing
	quit := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quit, os.Interrupt)

	go common.ShutdownServer(srv, quit, done)

	common.StartServer(&addrPtr, srv)

	<-done
	log.Println("Server gracefully stopped")
}

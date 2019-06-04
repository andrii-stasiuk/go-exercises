package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handler"
	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/router"

	//	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

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

	sql := model.New(dataBase)
	sqlVersion, err := sql.GetVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SQL Server version: %s\n", sqlVersion)

	srv := core.NewServer(addrPtr, router.NewRouter(router.AllRoutes(handler.New(&sql))))

	done := make(chan struct{}, 1)
	// Setting up signal capturing
	quit := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quit, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(quit, syscall.SIGTERM)

	go core.ShutdownServer(srv, quit, done)

	core.StartServer(addrPtr, srv)

	<-done
	fmt.Println("Server gracefully stopped")
}

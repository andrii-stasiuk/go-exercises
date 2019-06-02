package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

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

	fmt.Println("API server starting...")

	//dataBase, err := sql.Open("postgres", "testuser:testpass@tcp(localhost:5555)/testdb?sslmode=disable")
	dataBase, err := sql.Open("mysql", *dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	dataBase.SetMaxIdleConns(100)
	defer dataBase.Close()

	router := router.NewRouter(router.AllRoutes(handler.Handlers{SQL: model.Model{Db: dataBase}}))

	srv := &http.Server{
		Handler:      router,
		Addr:         *addrPtr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("API server started successfully on " + *addrPtr)
	log.Fatal(srv.ListenAndServe())
}

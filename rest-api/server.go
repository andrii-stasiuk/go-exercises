package main

import (
	"database/sql"
	"flag"
	"fmt"
	"go-exercises/rest-api/handler"
	"go-exercises/rest-api/model"
	"go-exercises/rest-api/router"
	"log"
	"net/http"
	"time"

	//	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var dbURLPtr = flag.String("db", "root:@tcp(127.0.0.1:3306)/testdb", "Specify the URL to the database")
	var addrPtr = flag.String("addr", "127.0.0.1:8000", "Server IPv4 address")
	flag.Parse()

	fmt.Println("API server starting...")

	//db, err := sql.Open("postgres", "testuser:testpass@tcp(localhost:5555)/testdb?sslmode=disable")
	dataBase, err := sql.Open("mysql", *dbURLPtr)
	if err != nil {
		log.Fatal(err)
	}
	dataBase.SetMaxIdleConns(100)
	defer dataBase.Close()

	ml := model.Model{Db: dataBase}
	hl := handler.Handlers{SQL: ml}

	// router := httprouter.New()
	router := router.NewRouter(router.AllRoutes(hl))

	srv := &http.Server{
		Handler: router,
		Addr:    *addrPtr,
		// Enforce timeouts for created servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("API server started successfully on " + *addrPtr)
	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"flag"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/andrii-stasiuk/go-exercises/client-server/database"
	"github.com/andrii-stasiuk/go-exercises/client-server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var newdbPtr = flag.Bool("create", false, "Create a new database")
	var addrPtr = flag.String("addr", "127.0.0.1:8000", "Server IPv4 address")
	var dbfilePtr = flag.String("file", "database.json", "Specify the name of the database file")
	flag.Parse()

	fmt.Println("API server starting...")

	var db *database.Database
	db = database.CreateDB(*dbfilePtr)

	if !(*newdbPtr) {
		db.LoadFromFile()
		fmt.Println("Loaded database: " + *dbfilePtr)
	} else {
		fmt.Println("Using database: " + *dbfilePtr)
	}

	hl := &handlers.Handlers{Database: *db}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", hl.DefaultHandler)
	r.HandleFunc("/users/", hl.UserCreator).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}/", hl.UserGetter).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/", hl.UserDeleter).Methods("DELETE")

	// Handlers reserved for testing purposes
	r.HandleFunc("/users/save/", hl.UserSaver).Methods("GET")
	r.HandleFunc("/users/load/", hl.UserLoader).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    *addrPtr,
		// Enforce timeouts for created servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("API server started successfully on " + *addrPtr)

	log.Fatal(srv.ListenAndServe())
}

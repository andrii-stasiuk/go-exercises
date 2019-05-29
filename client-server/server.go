package main

import (
	"flag"
	"fmt"
	"go-exercises/client-server/database"
	"go-exercises/client-server/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var newdbPtr = flag.Bool("create", false, "Create a new database")
	var addrPtr = flag.String("addr", "127.0.0.1:8000", "Server IPv4 address")
	var dbfilePtr = flag.String("file", "database.json", "Specify the name of the database file")
	// #1.1 var delayPtr = flag.Int("delay", 60, "Delay in seconds between saving the database"
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

	var hl *handlers.Handlers
	hl = &handlers.Handlers{Database: *db}

	// #1.2 This is done to reduce server load
	// if *delayPtr < 5 {
	// 	*delayPtr = 5
	// }
	// fmt.Printf("Autosave interval set to: %d seconds\n", *delayPtr)

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

	// #1.3. go db.RepeatSaving(*delayPtr)

	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"flag"
	"fmt"
	handlers "go-exercises/client-server/inc"
	io "go-exercises/client-server/inc"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var databasePtr = flag.Bool("database", false, "Database load at program startup")
	var delayPtr = flag.Int("delay", 60, "Delay in seconds between saving the database")
	var addrPtr = flag.String("addr", "127.0.0.1", "Server IPv4 address")
	flag.Parse()

	fmt.Println("API server starting...")

	if *databasePtr {
		io.LoadFromFile(io.DataFile)
	}
	fmt.Println("Loaded database: " + strconv.FormatBool(*databasePtr))

	safeDelay := 5
	if *delayPtr > safeDelay {
		safeDelay = *delayPtr
	}
	fmt.Printf("Autosave interval set to: %d seconds\n", safeDelay)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", handlers.DefaultHandler)
	r.HandleFunc("/users/", handlers.UserCreator).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}/", handlers.UserGetter).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/", handlers.UserDeleter).Methods("DELETE")

	// Handlers reserved for testing purposes
	r.HandleFunc("/users/save/", handlers.UserSaver).Methods("GET")
	r.HandleFunc("/users/load/", handlers.UserLoader).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    io.CheckIP(*addrPtr) + ":8000",
		// Enforce timeouts for created servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("API server started successfully on " + io.CheckIP(*addrPtr))

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(safeDelay))
			if io.SaveToFile(io.DataFile) {
				fmt.Println("The database was backed up at", time.Now())
			} else {
				fmt.Println("An unknown error occurred while database saving")
			}
		}
	}()

	log.Fatal(srv.ListenAndServe())
}

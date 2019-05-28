package main

import (
	"flag"
	"fmt"
	database "go-exercises/client-server/inc"
	handlers "go-exercises/client-server/inc"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CheckIP function is used to verify the correctness of the IP-address
// and setting default value in case of error
func CheckIP(addrFlag string) string {
	if net.ParseIP(addrFlag) != nil {
		return addrFlag
	}
	return "127.0.0.1" // Returns default IP-address of the localhost
}

func main() {
	var databasePtr = flag.Bool("database", false, "Database load at program startup")
	var delayPtr = flag.Int("delay", 60, "Delay in seconds between saving the database")
	var addrPtr = flag.String("addr", "127.0.0.1", "Server IPv4 address")
	flag.Parse()

	fmt.Println("API server starting...")

	if *databasePtr {
		database.GetDBInstance().LoadFromFile(database.DataFile)
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
		Addr:    CheckIP(*addrPtr) + ":8000",
		// Enforce timeouts for created servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("API server started successfully on " + CheckIP(*addrPtr))

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(safeDelay))
			if err := database.GetDBInstance().SaveToFile(database.DataFile); err == nil {
				fmt.Println("The database was backed up at", time.Now())
			} else {
				log.Println(err)
			}
		}
	}()

	log.Fatal(srv.ListenAndServe())
}

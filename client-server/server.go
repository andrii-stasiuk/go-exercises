package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP Error 503 (Service Unavailable)\n"))
}

func UserCreator(w http.ResponseWriter, r *http.Request) {
	mutex := new(sync.RWMutex)
	innerDB := database{*mutex, map[uint64]User{}}
	innerDB.Set()
	w.Write([]byte("UserCreator\n"))
}

func UserDeleter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Delete " + vars["id"]))
	//w.Write([]byte("UserDeleter\n"))
}

func UserGetter(w http.ResponseWriter, r *http.Request) {
	//s := User{1, "2", "3", "4"}
	vars := mux.Vars(r)
	//w.Write([]byte("Get " + vars["id"]))

	innerDB := new(database)
	//d.Get(12)
	strtoint, _ := strconv.ParseUint(vars["id"], 10, 64)
	//n, err := strconv.ParseInt(s, 10, 64)
	// if err == nil {
	// 	fmt.Printf("%d of type %T", n, n)
	// }

	w.Write([]byte(innerDB.Get(int(strtoint)).Name + innerDB.Get(int(strtoint)).Surname))
	//w.Write([]byte("UserGetter\n"))
}

type database struct {
	m     sync.RWMutex
	users map[uint64]User
}

func (r database) Delete() {
}

func (r database) Get(us int) (usr User) {
	if us == 1 {
		return User{1, "Andrii", "Stasiuk", "as@ges.sh"}
	} else {
		return r.users[1]
	}
}

func (r *database) Set() {
	r.users[1] = User{1, "Andrii2", "Stasiuk2", "as@ges.sh"}
	//r.users[id] := 12
}

//var innerDB database

func main() {
	fmt.Println("Apache Server Started")
	//	innerDB := new(database)
	//innerDB.Set()
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	r.HandleFunc("/users/", UserCreator).Methods("GET") // .Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}/", UserDeleter).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}/", UserGetter).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8001", r))
}

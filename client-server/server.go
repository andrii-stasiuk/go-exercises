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
	// mutex := new(sync.RWMutex)
	// innerDB := Database{*mutex, map[uint64]User{}}
	// innerDB.Set()
	w.Write([]byte("Created user: " + NewDB().Set().Name + " " + NewDB().Set().Surname))
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
	// mutex := new(sync.RWMutex)
	// innerDB := Database{*mutex, map[uint64]User{}}
	//innerDB := new(database)
	//d.Get(12)
	strtoint, _ := strconv.ParseUint(vars["id"], 10, 64)
	//n, err := strconv.ParseInt(s, 10, 64)
	// if err == nil {
	// 	fmt.Printf("%d of type %T", n, n)
	// }
	w.Write([]byte("Get user: " + NewDB().Get(int(strtoint)).Name + " " + NewDB().Get(int(strtoint)).Surname))
	//w.Write([]byte("UserGetter\n"))
}

type Database struct {
	m     sync.RWMutex
	users map[uint64]User
}

func (r *Database) Delete() {
	delete(r.users, 2)
}

func (r *Database) Get(us int) User {
	if us == 1 {
		return User{1, "Andrii", "Stasiuk", "as@ges.sh"}
	} else {
		return r.users[2]
	}
}

func (r *Database) Set() User {
	r.users[2] = User{2, "John", "Smith", "as@ges.sh"}
	return r.users[2]
	//r.users[id] := 12
}

var innerDB *Database

//var lock = &sync.Mutex{}
func NewDB() *Database {
	// lock.Lock()
	// defer lock.Unlock()
	if innerDB == nil {
		mutex := new(sync.RWMutex)
		innerDB = new(Database)
		innerDB = &Database{*mutex, map[uint64]User{}}
	}
	return innerDB
	// mutex := new(sync.RWMutex)
	// return &Database{*mutex, map[uint64]User{}}
}

func main() {
	fmt.Println("Apache Server Started")
	// mutex := new(sync.RWMutex)
	// innerDB := Database{*mutex, map[uint64]User{}}
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

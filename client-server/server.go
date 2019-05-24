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

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP Error 503 (Service Unavailable)\n"))
}

func UserCreator(w http.ResponseWriter, r *http.Request) {
	id := NewDB().Set().ID
	usr := NewDB().Get(id)
	w.Write([]byte("Created user: " + strconv.FormatUint(usr.ID, 10) + " " + usr.Name + " " + usr.Surname))
}

func UserDeleter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strtoint, _ := strconv.ParseUint(vars["id"], 10, 64)
	w.Write([]byte("Deleted: " + NewDB().Delete(strtoint)))
}

func UserGetter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strtoint, _ := strconv.ParseUint(vars["id"], 10, 64)
	//n, err := strconv.ParseInt(s, 10, 64)
	// if err == nil {
	// 	fmt.Printf("%d of type %T", n, n)
	// }
	w.Write([]byte("Get user: " + NewDB().Get(uint64(strtoint)).Name + " " + NewDB().Get(uint64(strtoint)).Surname))
}

type Database struct {
	m     sync.RWMutex
	users map[uint64]User
}

func (r *Database) Delete(usr uint64) string {
	dl := r.users[usr].Name + r.users[usr].Surname
	delete(r.users, usr)
	return dl
}

func (r *Database) Get(usr uint64) User {
	// if us == 1 {
	// 	return User{1, "Andrii", "Stasiuk", "as@ges.sh"}
	// } else {
	return r.users[usr]
	// }
}

func (r *Database) Set() User {
	lastElem++
	r.users[lastElem] = User{lastElem, "John" + strconv.FormatUint(lastElem, 10), "Smith", "as@ges.sh"}
	return r.users[lastElem]
	//r.users[id] := 12
}

var innerDB *Database
var lastElem uint64

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
}

func main() {
	fmt.Println("Apache Server Started")
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/", DefaultHandler)

	r.HandleFunc("/users/", UserCreator).Methods("GET")                 // .Methods("POST")
	r.HandleFunc("/users/del/{id:[0-9]+}/", UserDeleter).Methods("GET") // .Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}/", UserGetter).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8001", r))
}

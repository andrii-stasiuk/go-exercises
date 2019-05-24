package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func UserSaver(w http.ResponseWriter, r *http.Request) {
	SaveToFile()
	w.Write([]byte("Saved\n"))
}

func UserLoader(w http.ResponseWriter, r *http.Request) {
	w.Write(LoadFromFile())
}

func UserCreator(w http.ResponseWriter, r *http.Request) {
	usrset, ok := NewDB().Set()
	if ok == true {
		usr, ok := NewDB().Get(usrset.ID)
		if ok == true {
			w.Write([]byte("Created user: " + strconv.FormatUint(usr.ID, 10) + " " + usr.Name + " " + usr.Surname))
		} else {
			w.Write([]byte("Error while getting data of created user"))
		}
	} else {
		w.Write([]byte("Error while creating user"))
	}
}

func UserDeleter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		usrset, deleted := NewDB().Delete(strtoint)
		if deleted == true {
			w.Write([]byte("Deleted: " + usrset))
		} else {
			w.Write([]byte("Error while getting user: " + usrset))
		}
	}
}

func UserGetter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		usrset, ok := NewDB().Get(uint64(strtoint))
		if ok == true {
			w.Write([]byte("Get user: " + usrset.Name + " " + usrset.Surname))
		} else {
			w.Write([]byte("Error while getting user"))
		}
	}
}

type Database struct {
	m     sync.RWMutex
	users map[uint64]User
}

func (r *Database) Delete(usrID uint64) (string, bool) {
	deletedUser := ""
	if _, exists := r.users[usrID]; exists {
		deletedUser = r.users[usrID].Name + " " + r.users[usrID].Surname
		delete(r.users, usrID)
	} else {
		return "Error (user doesnt exists)", false
	}
	if _, exists := r.users[usrID]; exists {
		return "Error (cant delete user)", false
	} else {
		return deletedUser, true
	}
}

func (r *Database) Get(usr uint64) (User, bool) {
	if _, ok := r.users[usr]; ok {
		return r.users[usr], true
	} else {
		return User{}, false
	}
}

func (r *Database) Set() (User, bool) {
	lastElem++
	r.users[lastElem] = User{lastElem, "John" + strconv.FormatUint(lastElem, 10), "Smith" + strconv.FormatUint(lastElem, 10), "as@ges.sh"}
	if _, ok := r.users[lastElem]; ok {
		return r.users[lastElem], true
	} else {
		return User{}, false
	}
}

func SaveToFile() {
	//NewDB().users[0] = User{0, "System", "Field", strconv.FormatUint(lastElem, 10)}
	//var us map[uint64]map[uint64]User
	us := make(map[uint64]map[uint64]User)
	us[lastElem] = NewDB().users

	//content, _ := json.Marshal(NewDB().users)
	content, _ := json.Marshal(us)
	ioutil.WriteFile("database.txt", []byte(content), 0644)
}

func LoadFromFile() []byte {
	content, _ := ioutil.ReadFile("database.txt")

	us := make(map[uint64]map[uint64]User) //
	json.Unmarshal([]byte(content), &us)   //
	// json.Unmarshal([]byte(content), &NewDB().users)

	for k := range us {
		lastElem = k
		NewDB().users = us[k]
	}

	// lastElem, _ = strconv.ParseUint(NewDB().users[0].Email, 10, 64)
	return content //[]byte(strconv.FormatUint(lastElem, 10))
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

	r.HandleFunc("/users/save/", UserSaver).Methods("GET")
	r.HandleFunc("/users/load/", UserLoader).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8002", r))
}

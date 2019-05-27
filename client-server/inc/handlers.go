package inc

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DefaultHandler function is used to process requests to the root path
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"Alive": true}`)
}

// UserCreator function is used to process user creation or modification requests
func UserCreator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	usrset, created := GetDBInstance().Set(r.FormValue("id"), r.FormValue("name"), r.FormValue("surname"), r.FormValue("email"))
	if created {
		respond, err := json.Marshal(usrset)
		if err != nil {
			panic(err)
		} else {
			w.Write(respond)
		}
	} else {
		io.WriteString(w, `{"Error": "Can't create/change user"}`)
	}
}

// UserGetter function is used to process requests for receiving user data
func UserGetter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		usrget, exists := GetDBInstance().Get(uint64(strtoint))
		if exists {
			respond, err := json.Marshal(usrget)
			if err != nil {
				panic(err)
			} else {
				w.Write(respond)
			}
		} else {
			respond, err := json.Marshal(map[string]string{"Error": "Can't get user with ID " + vars["id"]})
			if err != nil {
				panic(err)
			} else {
				w.Write(respond)
			}
		}
	} else {
		panic(err)
	}
}

// UserDeleter function is used to process user deletion requests
func UserDeleter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		if usrdel, deleted := GetDBInstance().Delete(strtoint); deleted {
			respond, err := json.Marshal(map[string]string{"Deleted user " + usrdel + " with ID": vars["id"]})
			if err != nil {
				panic(err)
			} else {
				w.Write(respond)
			}
		} else {
			respond, err := json.Marshal(map[string]string{"Error": "Can't delete user with ID " + vars["id"]})
			if err != nil {
				panic(err)
			} else {
				w.Write(respond)
			}
		}
	} else {
		panic(err)
	}
}

// UserSaver function is used to process requests to save the user database
func UserSaver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := SaveToFile(DataFile); err == nil {
		io.WriteString(w, `{"Status": "Database saved"}`)
	} else {
		io.WriteString(w, `{"Error": "Can't save databcdase to a file"}`)
		log.Println(err)
	}
}

// UserLoader function is used to process requests to load the user database
func UserLoader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if data, loaded := LoadFromFile(DataFile); loaded {
		w.Write(data)
	}
}

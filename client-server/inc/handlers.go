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
	usrset, err := GetDBInstance().Set(r.FormValue("id"), r.FormValue("name"), r.FormValue("surname"), r.FormValue("email"))
	if err == nil {
		respond, err := json.Marshal(usrset)
		if err != nil {
			log.Println(err)
			io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
		} else {
			w.Write(respond)
		}
	} else {
		log.Println(err)
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
		usrget, err := GetDBInstance().Get(uint64(strtoint))
		if err == nil {
			respond, err := json.Marshal(usrget)
			if err != nil {
				log.Println(err)
				io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
			} else {
				w.Write(respond)
			}
		} else {
			log.Println(err)
			respond, err := json.Marshal(map[string]string{"Error": "Can't get user with ID " + vars["id"]})
			if err != nil {
				log.Println(err)
				io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
			} else {
				w.Write(respond)
			}
		}
	} else {
		io.WriteString(w, `{"Error": "Can't convert string field ID to integer"}`)
		log.Println(err)
	}
}

// UserDeleter function is used to process user deletion requests
func UserDeleter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		if usrdel, err := GetDBInstance().Delete(strtoint); err == nil {
			respond, err := json.Marshal(map[string]string{"Deleted user " + usrdel + " with ID": vars["id"]})
			if err != nil {
				log.Println(err)
				io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
			} else {
				w.Write(respond)
			}
		} else {
			log.Println(err)
			respond, err := json.Marshal(map[string]string{"Error": "Can't delete user with ID " + vars["id"]})
			if err != nil {
				log.Println(err)
				io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
			} else {
				w.Write(respond)
			}
		}
	} else {
		io.WriteString(w, `{"Error": "Can't convert string field ID to integer"}`)
		log.Println(err)
	}
}

// UserSaver function is used to process requests to save the user database
func UserSaver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := GetDBInstance().SaveToFile(DataFile); err == nil {
		io.WriteString(w, `{"Status": "Database saved"}`)
	} else {
		io.WriteString(w, `{"Error": "Can't save database to a file"}`)
		log.Println(err)
	}
}

// UserLoader function is used to process requests to load the user database
func UserLoader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if data, err := GetDBInstance().LoadFromFile(DataFile); err == nil {
		w.Write(data)
	} else {
		io.WriteString(w, `{"Error": "Can't load database from a file"}`)
		log.Println(err)
	}
}

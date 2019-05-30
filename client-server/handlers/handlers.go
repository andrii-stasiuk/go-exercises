/*Package handlers Controller*/
package handlers

import (
	"encoding/json"

	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/andrii-stasiuk/go-exercises/client-server/database"
	"github.com/gorilla/mux"
)

// Handlers structure
type Handlers struct {
	Database database.Database
}

// DefaultHandler function is used to process requests to the root path
func (hl *Handlers) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"Alive": true}`)
}

// UserCreator function is used to process user creation or modification requests
func (hl *Handlers) UserCreator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	usrset, err := hl.Database.Set(r.FormValue("id"), r.FormValue("name"), r.FormValue("surname"), r.FormValue("email"))
	if err == nil {
		respond, err := json.Marshal(usrset)
		if err != nil {
			log.Println(err)
			io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
		} else {
			w.Write(respond)
		}
		// Saving the database when some user was actually added or changed
		if err := hl.Database.SaveToFile(); err == nil {
			// io.WriteString(w, `{"Status": "Database saved"}`)
		} else {
			io.WriteString(w, `{"Error": "Can't save database to a file"}`)
			log.Println(err)
		}
	} else {
		log.Println(err)
		io.WriteString(w, `{"Error": "Can't create/change user"}`)
	}
}

// UserGetter function is used to process requests for receiving user data
func (hl *Handlers) UserGetter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		usrget, err := hl.Database.Get(uint64(strtoint))
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
func (hl *Handlers) UserDeleter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	strtoint, err := strconv.ParseUint(vars["id"], 10, 64)
	if err == nil {
		if usrdel, err := hl.Database.Delete(strtoint); err == nil {
			respond, err := json.Marshal(map[string]string{"Deleted user " + usrdel + " with ID": vars["id"]})
			if err != nil {
				log.Println(err)
				io.WriteString(w, `{"Error": "Can't encode into JSON"}`)
			} else {
				w.Write(respond)
			}
			// Saving the database when some user was actually deleted
			if err := hl.Database.SaveToFile(); err == nil {
				// io.WriteString(w, `{"Status": "Database saved"}`)
			} else {
				io.WriteString(w, `{"Error": "Can't save database to a file"}`)
				log.Println(err)
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
func (hl *Handlers) UserSaver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := hl.Database.SaveToFile(); err == nil {
		io.WriteString(w, `{"Status": "Database saved"}`)
	} else {
		io.WriteString(w, `{"Error": "Can't save database to a file"}`)
		log.Println(err)
	}
}

// UserLoader function is used to process requests to load the user database
func (hl *Handlers) UserLoader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if data, err := hl.Database.LoadFromFile(); err == nil {
		w.Write(data)
	} else {
		io.WriteString(w, `{"Error": "Can't load database from a file"}`)
		log.Println(err)
	}
}

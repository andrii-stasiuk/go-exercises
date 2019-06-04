/*Package handler Todo*/
package handler

import (
	"encoding/json"

	"log"
	"net/http"
	"strconv"

	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// Handlers structure
type Handlers struct {
	SQL model.Model
}

// Default - handler for the root page
func (h *Handlers) Default(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// w.Header().Set("Content-Type", "text/html")
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w, "Welcome!\n")
	responses.WriteOKResponse(w, "Welcome to API server!")
}

// TodoIndex - handler for the Todo Index action
func (h *Handlers) TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if res, err := h.SQL.Index(); err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		responses.WriteOKResponse(w, res)
	}
}

// TodoShow - handler for the Todo Show action
func (h *Handlers) TodoShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
	} else {
		res, err := h.SQL.Show(params.ByName("id"))
		if err != nil {
			log.Println(err)
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		} else {
			responses.WriteOKResponse(w, res)
		}
	}
}

// TodoDelete - handler for the Todo Delete action
func (h *Handlers) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
	} else {
		deleted, err := h.SQL.Delete(params.ByName("id"))
		if err != nil {
			log.Println(err)
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		} else if deleted == true {
			responses.WriteOKResponse(w, http.StatusOK)
		} else {
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		}
	}
}

// TodoCreate - handler for the Todo Create action
func (h *Handlers) TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		if !h.CheckStr(todo.Name) || !h.CheckStr(todo.Description) || !h.CheckInt(todo.State) {
			log.Println("Incorrect input data")
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		} else {
			res, err := h.SQL.Create(&todo)
			if err != nil {
				log.Println(err)
				responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
			} else {
				responses.WriteOKResponse(w, res)
			}
		}

	}
}

// TodoUpdate - handler for the Todo Update action
func (h *Handlers) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		if !h.CheckInt(params.ByName("id")) || !h.CheckStr(todo.Name) || !h.CheckStr(todo.Description) || !h.CheckInt(todo.State) {
			log.Println("Incorrect input data")
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		} else {
			res, err := h.SQL.Update(params.ByName("id"), &todo)
			if err != nil {
				log.Println(err)
				responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
			} else {
				responses.WriteOKResponse(w, res)
			}
		}
	}
}

// CheckInt basic check of integer
func (h *Handlers) CheckInt(id string) bool {
	converted, err := strconv.ParseUint(id, 10, 64)
	if err == nil && converted > 0 {
		return true
	}
	return false
}

// CheckStr basic check of string
func (h *Handlers) CheckStr(str string) bool {
	if len(str) > 0 && str != "`" {
		return true
	}
	return false
}
package handler

import (
	"go-exercises/rest-api/errors"
	"go-exercises/rest-api/model"
	"log"
	"net/http"

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
	errors.WriteOKResponse(w, "Welcome to API server!")
}

// TodoIndex - handler for the Todo Index action
func (h *Handlers) TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if res, err := h.SQL.Index(); err != nil {
		log.Println(err)
		errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		errors.WriteOKResponse(w, res)
	}
}

// TodoCreate - handler for the Todo Create action
func (h *Handlers) TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		res, err := h.SQL.Create(r.Form.Get("name"), r.Form.Get("description"), r.Form.Get("state"))
		if err != nil {
			log.Println(err)
			errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		} else {
			errors.WriteOKResponse(w, res)
		}
	}
}

// TodoShow - handler for the Todo Show action
func (h *Handlers) TodoShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res, err := h.SQL.Show(params.ByName("id"))
	if err != nil {
		log.Println(err)
		errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		errors.WriteOKResponse(w, res)
	}
}

// TodoUpdate - handler for the Todo Update action
func (h *Handlers) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		res, err := h.SQL.Update(params.ByName("id"), r.Form.Get("name"), r.Form.Get("description"), r.Form.Get("state"))
		if err != nil {
			log.Println(err)
			errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		} else {
			errors.WriteOKResponse(w, res)
		}
	}
	// jsonResponse(w, h.SQL.Update(t, r.Form.Get("name"), r.Form.Get("description")))
}

// TodoDelete - handler for the Todo Delete action
func (h *Handlers) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := h.SQL.Delete(params.ByName("id"))
	if err != nil {
		log.Println(err)
		errors.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
	} else {
		errors.WriteOKResponse(w, http.StatusOK)
	}
	//w.WriteHeader(200)
}

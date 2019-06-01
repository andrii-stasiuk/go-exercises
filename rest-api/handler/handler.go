package handler

import (
	"encoding/json"
	"fmt"
	"go-exercises/rest-api/errors"
	"go-exercises/rest-api/model"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Handlers structure
type Handlers struct {
	Database model.Model
}

func (s *Handlers) Default(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (s *Handlers) TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonResponse(w, s.Database.Index())
}

func (s *Handlers) TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	jsonResponse(w, s.Database.Create(r.Form.Get("name"), r.Form.Get("description")))
}

func (s *Handlers) TodoShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	t, _ := strconv.ParseUint(id, 10, 64)
	jsonResponse(w, s.Database.Show(t))
}

func (s *Handlers) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	t, _ := strconv.ParseUint(id, 10, 64)
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	jsonResponse(w, s.Database.Update(t, r.Form.Get("name"), r.Form.Get("description")))
}

func (s *Handlers) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	t, _ := strconv.ParseUint(id, 10, 64)
	s.Database.Delete(t)
	w.WriteHeader(200)
}

func jsonResponse(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	if errors.ErrorCheck(res, err) {
		return
	}

	fmt.Fprintf(res, string(payload))
}

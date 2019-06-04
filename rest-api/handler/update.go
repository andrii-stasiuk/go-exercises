package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

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

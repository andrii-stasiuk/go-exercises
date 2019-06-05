package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoUpdate - handler for the Todo Update action
func (h Handlers) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	if !CheckInt(params.ByName("id")) || !CheckStr(todo.Name) || !CheckStr(todo.Description) || !CheckInt(todo.State) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	res, err := h.SQL.Update(params.ByName("id"), &todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

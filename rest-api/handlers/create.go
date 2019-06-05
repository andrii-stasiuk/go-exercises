package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoCreate - handler for the Todo Create action
func (h Handlers) TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	if !CheckStr(todo.Name) || !CheckStr(todo.Description) || !CheckInt(todo.State) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	res, err := h.SQL.Create(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

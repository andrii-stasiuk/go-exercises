/*Package handlers Todo*/
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoUpdate - handler for the Todo Update action, also validates the data received from the client
func (h Handlers) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := model.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !CheckInt(params.ByName("id")) || !CheckStr(todo.Name) || !CheckStr(todo.Description) || !CheckInt(todo.State) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	id64, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
		return
	}
	todo.ID = int(id64)
	res, err := h.SQL.Update(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

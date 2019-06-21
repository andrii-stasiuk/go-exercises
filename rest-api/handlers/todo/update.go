/*Package handlers Todo*/
package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/todomodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoUpdate - handler for the Todo Update action, also validates the data received from the client
func (h TodoHandlers) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := todomodel.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !core.CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
		return
	}
	if !(core.CheckStr(todo.Name) || core.CheckStr(todo.Description) || core.CheckInt(todo.State)) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	todo.ID, err = strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
		return
	}
	res, err := h.SQL.Update(todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

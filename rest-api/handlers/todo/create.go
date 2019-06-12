/*Package handlers Todo*/
package todo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/todomodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoCreate - handler for the Todo Create action, also validates the data received from the client
func (h TodoHandlers) TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := todomodel.Todo{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !core.CheckStr(todo.Name) || !core.CheckStr(todo.Description) || !core.CheckInt(todo.State) {
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

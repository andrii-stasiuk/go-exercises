/*Package handlers Todo*/
package todohandlers

import (
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoDelete - handler for the Todo Delete action, also validates the "id" field
func (h TodoHandlers) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !core.CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
		return
	}
	res, err := h.SQL.Delete(params.ByName("id"))
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

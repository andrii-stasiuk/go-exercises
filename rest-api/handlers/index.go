package handlers

import (
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoIndex - handler for the Todo Index action
func (h Handlers) TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.SQL.Index()
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

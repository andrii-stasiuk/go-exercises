package handlers

import (
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoShow - handler for the Todo Show action
func (h Handlers) TodoShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
		return
	}
	res, err := h.SQL.Show(params.ByName("id"))
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

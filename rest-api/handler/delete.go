package handler

import (
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoDelete - handler for the Todo Delete action
func (h *Handlers) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
	} else {
		deleted, err := h.SQL.Delete(params.ByName("id"))
		if err != nil {
			log.Println(err)
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		} else if deleted == true {
			responses.WriteOKResponse(w, http.StatusOK)
		} else {
			responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		}
	}
}

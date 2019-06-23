/*Package handlers Todo*/
package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoIndex - handler for the Todo Index action
func (h TodoHandlers) TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Context().Value("user_id").(uint64)
	if userID < 1 {
		log.Println("Incorrect User ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect User ID")
		return
	}
	res, err := h.SQL.Index(strconv.FormatUint(userID, 10))
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

/*Package handlers Todo*/
package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/andrii-stasiuk/go-exercises/rest-api/common"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// TodoDelete - handler for the Todo Delete action, also validates the "id" field
func (h TodoHandlers) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := r.Context().Value("user_id").(uint64)
	if userID < 1 {
		log.Println("Incorrect User ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect User ID")
		return
	}
	if !common.CheckInt(params.ByName("id")) {
		log.Println("Incorrect ID")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect ID")
		return
	}
	res, err := h.SQL.Delete(params.ByName("id"), strconv.FormatUint(userID, 10))
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	responses.WriteOKResponse(w, res)
}

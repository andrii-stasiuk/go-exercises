/*Package handlers Todo*/
package todo

import (
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// Default - handler for the root page /
func (h TodoHandlers) Default(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responses.WriteOKResponse(w, "Welcome to API server!")
}

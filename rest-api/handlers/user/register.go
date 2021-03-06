/*Package userhandler*/
package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/validemail"

	"github.com/andrii-stasiuk/go-exercises/rest-api/common"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/julienschmidt/httprouter"
)

// UserRegister - handler for the user register action
func (uh UserHandlers) UserRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := usermodel.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !validemail.New().EMailValidator(user.Email) || !common.CheckStr(user.Password) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	res, err := uh.SQL.Register(user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Can't create new user")
		return
	}
	responses.WriteOKResponse(w, res)
}

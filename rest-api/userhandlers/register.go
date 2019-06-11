/*Package userhandler*/
package userhandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/validemail"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/andrii-stasiuk/go-exercises/rest-api/usermodel"
	"github.com/julienschmidt/httprouter"
)

//
func (uh UserHandlers) UserRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := usermodel.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !validemail.EMailValidator(user.Email) || !core.CheckStr(user.Password) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	user.Password, err = core.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Can't hash the password")
		return
	}
	res, err := uh.SQL.Register(&user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Can't create new user")
		return
	}
	responses.WriteOKResponse(w, res)
}

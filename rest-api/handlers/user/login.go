/*Package userhandler*/
package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/auth"
	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/andrii-stasiuk/validemail"
	"github.com/julienschmidt/httprouter"
)

// UserLogin - handler for the user login action
func (uh UserHandlers) UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := usermodel.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !validemail.New().EMailValidator(user.Email) || !core.CheckStr(user.Password) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	res, logged := uh.SQL.Login(user)
	if !logged {
		log.Println("Unable to login")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to login")
		return
	}
	tokenString, expirationTime := auth.GetToken(res)
	w.Header().Set("Authorization", "Bearer "+tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "x-access-token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	resp := make(map[string]interface{})
	resp["token"] = tokenString // Store the token in the response
	resp["user"] = map[string]interface{}{"id": res.ID, "email": res.Email, "created_at": res.CreatedAt}
	responses.WriteOKResponse(w, resp)
}

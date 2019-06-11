/*Package userhandler*/
package userhandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrii-stasiuk/go-exercises/rest-api/core"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers"
	"github.com/andrii-stasiuk/go-exercises/rest-api/responses"
	"github.com/andrii-stasiuk/go-exercises/rest-api/usermodel"
	"github.com/julienschmidt/httprouter"
)

type userInterface interface {
	Register(user *usermodel.User) (*usermodel.User, error)
	Login(user *usermodel.User) bool
}

// Handlers structure for handling requests
type UserHandler struct {
	SQL userInterface
}

// New is a constructor ...
func New(us userInterface) UserHandler {
	return UserHandler{SQL: us}
}

func (uh UserHandler) UserRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := usermodel.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !handlers.CheckStr(user.Email) || !handlers.CheckStr(user.Password) {
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

func (uh UserHandler) UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := usermodel.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to decode JSON")
		return
	}
	if !handlers.CheckStr(user.Email) || !handlers.CheckStr(user.Password) {
		log.Println("Incorrect input data")
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Incorrect input data")
		return
	}
	// user.Password, err = core.HashPassword(user.Password)
	// if err != nil {
	// 	log.Println(err)
	// 	responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Can't hash the password")
	// 	return
	// }
	res := uh.SQL.Login(&user)
	if res == false {
		//log.Println(err)
		responses.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Unable to login")
		return
	}
	responses.WriteOKResponse(w, res)
}

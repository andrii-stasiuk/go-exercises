/*Package userhandler*/
package user

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/usermodel"
)

type userInterface interface {
	Register(user usermodel.User) (usermodel.User, error)
	Login(user usermodel.User) (usermodel.User, bool)
}

// Handlers structure for handling requests
type UserHandlers struct {
	SQL userInterface
}

// New is a constructor ...
func New(us userInterface) UserHandlers {
	return UserHandlers{SQL: us}
}

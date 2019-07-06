/*Package router - user routes*/
package router

import "github.com/andrii-stasiuk/go-exercises/rest-api/handlers/user"

// UserRoutes - Routes for Users
func UserRoutes(uh user.UserHandlers) Routes {
	return Routes{
		Route{"POST", "/api/users/", false, uh.UserRegister},
		Route{"POST", "/api/users/login/", false, uh.UserLogin},
	}
}

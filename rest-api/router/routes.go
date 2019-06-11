/*Package router (all routes)*/
package router

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers"
	"github.com/andrii-stasiuk/go-exercises/rest-api/userhandler"
)

// AllRoutes is here
func AllRoutes(hl handlers.Handlers, us userhandler.UserHandler) Routes {
	routes := Routes{
		Route{"GET", "/", hl.Default},
		Route{"GET", "/api/todos/", hl.TodoIndex},
		Route{"GET", "/api/todos/:id/", hl.TodoShow},
		Route{"DELETE", "/api/todos/:id/", hl.TodoDelete},
		Route{"POST", "/api/todos/", hl.TodoCreate},
		Route{"PATCH", "/api/todos/:id/", hl.TodoUpdate},

		Route{"POST", "/api/user/register/", us.UserRegister},
		Route{"POST", "/api/user/login/", us.UserLogin},
	}
	return routes
}

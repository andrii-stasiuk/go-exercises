/*Package router (all routes)*/
package router

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/todohandlers"
	"github.com/andrii-stasiuk/go-exercises/rest-api/userhandlers"
)

// AllRoutes is here
func AllRoutes(hl todohandlers.TodoHandlers, us userhandlers.UserHandlers) Routes {
	routes := Routes{
		// Routes for Todos
		Route{"GET", "/", hl.Default},
		Route{"GET", "/api/todos/", hl.TodoIndex},
		Route{"GET", "/api/todos/:id/", hl.TodoShow},
		Route{"DELETE", "/api/todos/:id/", hl.TodoDelete},
		Route{"POST", "/api/todos/", hl.TodoCreate},
		Route{"PATCH", "/api/todos/:id/", hl.TodoUpdate},
		// Routes for Users
		Route{"POST", "/api/user/register/", us.UserRegister},
		Route{"POST", "/api/user/login/", us.UserLogin},
	}
	return routes
}

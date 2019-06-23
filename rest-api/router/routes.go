/*Package router (all routes)*/
package router

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers/todo"
	"github.com/andrii-stasiuk/go-exercises/rest-api/handlers/user"
)

// TodoRoutes - Routes for Todos
func TodoRoutes(hl todo.TodoHandlers) Routes {
	routes := Routes{
		Route{"GET", "/", false, hl.Default},
		Route{"GET", "/api/todos/", true, hl.TodoIndex},
		Route{"GET", "/api/todos/:id/", true, hl.TodoShow},
		Route{"DELETE", "/api/todos/:id/", true, hl.TodoDelete},
		Route{"POST", "/api/todos/", true, hl.TodoCreate},
		Route{"PATCH", "/api/todos/:id/", true, hl.TodoUpdate},
	}
	return routes
}

// UserRoutes - Routes for Users
func UserRoutes(us user.UserHandlers) Routes {
	routes := Routes{
		Route{"POST", "/api/users/", false, us.UserRegister},
		Route{"POST", "/api/users/login/", false, us.UserLogin},
	}
	return routes
}

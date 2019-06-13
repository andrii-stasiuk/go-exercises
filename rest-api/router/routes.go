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
		Route{"GET", "/api/todos/", false, hl.TodoIndex},
		Route{"GET", "/api/todos/:id/", false, hl.TodoShow},
		Route{"DELETE", "/api/todos/:id/", false, hl.TodoDelete},
		Route{"POST", "/api/todos/", false, hl.TodoCreate},
		Route{"PATCH", "/api/todos/:id/", false, hl.TodoUpdate},
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

// ProtectedRoutes - Routes that require Authentication
func ProtectedRoutes(hl todo.TodoHandlers) Routes {
	routes := Routes{
		Route{"POST", "/api/test/", true, hl.Default}, // temporary, for development purposes
	}
	return routes
}

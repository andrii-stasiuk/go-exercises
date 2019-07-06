/*Package router - todo routes*/
package router

import "github.com/andrii-stasiuk/go-exercises/rest-api/handlers/todo"

// TodoRoutes - Routes for Todos
func TodoRoutes(th todo.TodoHandlers) Routes {
	return Routes{
		Route{"GET", "/", false, th.Default},
		Route{"GET", "/api/todos/", true, th.TodoIndex},
		Route{"GET", "/api/todos/:id/", true, th.TodoShow},
		Route{"DELETE", "/api/todos/:id/", true, th.TodoDelete},
		Route{"POST", "/api/todos/", true, th.TodoCreate},
		Route{"PATCH", "/api/todos/:id/", true, th.TodoUpdate},
	}
}

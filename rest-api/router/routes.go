/*Package router*/
package router

import "github.com/andrii-stasiuk/go-exercises/rest-api/handlers"

// AllRoutes
func AllRoutes(hl handlers.Handlers) Routes {
	routes := Routes{
		Route{"GET", "/", hl.Default},
		Route{"GET", "/api/todos/", hl.TodoIndex},
		Route{"GET", "/api/todos/:id/", hl.TodoShow},
		Route{"DELETE", "/api/todos/:id/", hl.TodoDelete},
		Route{"POST", "/api/todos/", hl.TodoCreate},
		Route{"PATCH", "/api/todos/:id/", hl.TodoUpdate},
	}
	return routes
}

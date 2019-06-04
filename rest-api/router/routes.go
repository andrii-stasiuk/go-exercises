/*Package router*/
package router

import "github.com/andrii-stasiuk/go-exercises/rest-api/handler"

// AllRoutes
func AllRoutes(hl handler.Handlers) Routes {
	routes := Routes{
		Route{"Default", "GET", "/", hl.Default},
		Route{"TodoIndex", "GET", "/api/todos/", hl.TodoIndex},
		Route{"TodoShow", "GET", "/api/todos/:id/", hl.TodoShow},
		Route{"TodoDelete", "DELETE", "/api/todos/:id/", hl.TodoDelete},
		Route{"TodoCreate", "POST", "/api/todos/", hl.TodoCreate},
		Route{"TodoUpdate", "PATCH", "/api/todos/:id/", hl.TodoUpdate},
	}
	return routes
}

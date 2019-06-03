/*Package router*/
package router

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/handler"
	"github.com/andrii-stasiuk/go-exercises/rest-api/logger"
	"github.com/julienschmidt/httprouter"
)

// Route - A new Route entry passed to the routes slice will be automatically
// translated to a handler with the NewRouter() function
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// Routes
type Routes []Route

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

// NewRouter - Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle
		handle = route.HandlerFunc
		handle = logger.Logger(handle)
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}

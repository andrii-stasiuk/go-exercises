package router

import (
	"go-exercises/rest-api/handler"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

/*
Define all the routes here.
A new Route entry passed to the routes slice will be automatically
translated to a handler with the NewRouter() function
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes(hl handler.Handlers) Routes {
	routes := Routes{
		Route{"Default", "GET", "/", hl.Default},
		Route{"TodoIndex", "GET", "/api/todos/", hl.TodoIndex},
		Route{"TodoCreate", "POST", "/api/todos/", hl.TodoCreate},
		Route{"TodoShow", "GET", "/api/todos/:id/", hl.TodoShow},
		Route{"TodoUpdate", "PATCH", "/api/todos/:id/", hl.TodoUpdate},
		Route{"TodoDelete", "DELETE", "/api/todos/:id/", hl.TodoDelete},
	}
	return routes
}

//Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle
		handle = route.HandlerFunc
		handle = Logger(handle)
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}

// A Logger function which simply wraps the handler function around some log messages
func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		fn(w, r, param)
		log.Printf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}

/*Package router (logic)*/
package router

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/logger"
	"github.com/julienschmidt/httprouter"
)

// Route - a new Route entry passed to the routes slice, will be automatically
// translated to a handler with the NewRouter() function
type Route struct {
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// Routes slice of Route
type Routes []Route

// NewRouter - reads from the routes slice to translate the values to httprouter.Handle
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

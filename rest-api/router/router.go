/*Package router (logic)*/
package router

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/auth"
	"github.com/andrii-stasiuk/go-exercises/rest-api/logger"
	"github.com/julienschmidt/httprouter"
)

// Route - a new Route entry passed to the routes slice, will be automatically
// translated to a handler with the NewRouter() function
type Route struct {
	Method      string
	Path        string
	Secured     bool
	HandlerFunc httprouter.Handle
}

// Routes slice of Route
type Routes []Route

// NewRouter - reads from the routes slice (of slices) to translate the values to httprouter.Handle
func NewRouter(routes ...Routes) *httprouter.Router {
	var ApplicationRoutes []Route
	// Converting a two-dimensional slice into a one-dimensional slice
	for _, routesElem := range routes {
		ApplicationRoutes = append(ApplicationRoutes, routesElem...)
	}
	router := httprouter.New()
	for _, route := range ApplicationRoutes {
		var handle httprouter.Handle
		handle = route.HandlerFunc
		handle = logger.Logger(handle)
		if route.Secured {
			handle = auth.Auth(handle)
		}
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}

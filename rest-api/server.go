package main

import (
	"database/sql"
	"fmt"
	"go-exercises/rest-api/handler"
	"go-exercises/rest-api/model"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	//	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

// A Logger function which simply wraps the handler function around some log messages
func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		fn(w, r, param)
		log.Printf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}

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

func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", Index},
		Route{"BookIndex", "GET", "/books", Index},
		Route{"Bookshow", "GET", "/books/:isdn", Index},
		Route{"Bookshow", "POST", "/books", Index},
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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func main() {
	//db, err := sql.Open("postgres", "testuser:testpass@tcp(localhost:5555)/testdb?sslmode=disable")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(100)
	defer db.Close()

	ml := model.Model{Db: db}
	hl := handler.Handlers{Database: ml}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/api/todos/", hl.TodoIndex)
	router.POST("/api/todos/", hl.TodoCreate)
	router.GET("/api/todos/:id/", hl.TodoShow)
	router.PATCH("/api/todos/:id/", hl.TodoUpdate)
	router.DELETE("/api/todos/:id/", hl.TodoDelete)

	// router := NewRouter(AllRoutes())
	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

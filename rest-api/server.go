package main

import (
	"database/sql"
	"fmt"
	"go-exercises/rest-api/handler"
	"go-exercises/rest-api/model"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "testuser:testpass@tcp(localhost:5555)/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(100)
	defer db.Close()

	server := &model.Server{Db: db}

	reHandler := new(handler.RegexpHandler)

	reHandler.HandleFunc("/api/todos/$", "POST", server.TodoCreate)
	reHandler.HandleFunc("/api/todos/$", "GET", server.TodoIndex)
	reHandler.HandleFunc("/api/todos/[0-9]+$", "GET", server.TodoShow)
	reHandler.HandleFunc("/api/todos/[0-9]+$", "PATCH", server.TodoUpdate)
	reHandler.HandleFunc("/api/todos/[0-9]+$", "DELETE", server.TodoDelete)
	//reHandler.HandleFunc(".*.[js|css|png|eof|svg|ttf|woff]", "GET", server.Assets)
	//reHandler.HandleFunc("/", "GET", server.Homepage)

	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", reHandler)
}

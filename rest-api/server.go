package main

import (
	"database/sql"
	"fmt"
	"go-exercises/rest-api/handler"
	"go-exercises/rest-api/model"
	"go-exercises/rest-api/router"
	"log"
	"net/http"

	//	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

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
	// router := httprouter.New()
	router := router.NewRouter(router.AllRoutes(hl))

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

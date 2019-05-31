package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// store "context" values and connections in the server struct
type Server struct {
	Db *sql.DB
}

// todo "Object"
type Todo struct {
	Id          int    `json:"id,sting"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"State"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func (s *Server) TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var todos []*Todo

	rows, err := s.Db.Query("SELECT id, name, description, state FROM todos")
	errorCheck(w, err)
	for rows.Next() {
		todo := &Todo{}
		rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.State)
		todos = append(todos, todo)
	}
	rows.Close()

	jsonResponse(w, todos)
}

func (s *Server) TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := &Todo{}

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	todo.Name = r.Form.Get("name")
	todo.Description = r.Form.Get("description")

	result, err := s.Db.Exec("INSERT INTO todos(name, description) VALUES(?, ?)", todo.Name, todo.Description)
	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	Id64, err := result.LastInsertId()
	Id := int(Id64)
	todo = &Todo{Id: Id}

	s.Db.QueryRow("SELECT state, name, description FROM todos WHERE Id=?", todo.Id).Scan(&todo.State, &todo.Name, &todo.Description)

	jsonResponse(w, todo)
}

func (s *Server) TodoShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	todo := &Todo{}
	err := s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id).Scan(&todo.Id, &todo.State, &todo.Name, &todo.Description, &todo.Created_at, &todo.Updated_at)
	if err != nil {
		fmt.Println("ERROR reading from db - ", err)
	}
	jsonResponse(w, todo)
}

func (s *Server) TodoUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todo := &Todo{}

	id := params.ByName("id")

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	todo.Name = r.Form.Get("name")
	todo.Description = r.Form.Get("description")

	_, err = s.Db.Exec("UPDATE todos SET name = ?, description = ? WHERE id = ?", todo.Name, todo.Description, id)

	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	err = s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id).Scan(&todo.Id, &todo.State, &todo.Name, &todo.Description, &todo.Created_at, &todo.Updated_at)
	if err != nil {
		fmt.Println("ERROR reading from db - ", err)
	}

	jsonResponse(w, todo)
}

func (s *Server) TodoDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	s.Db.Exec("DELETE FROM todos WHERE id=?", id)
	w.WriteHeader(200)
}

func jsonResponse(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	if errorCheck(res, err) {
		return
	}

	fmt.Fprintf(res, string(payload))
}

func errorCheck(res http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

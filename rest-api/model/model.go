package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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

func (s *Server) Homepage(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Content-Type", "application/json")
	// res.WriteHeader(http.StatusOK)
	// io.WriteString(res, `{"Alive": true}`)
	comm := map[string]string{
		"rsc": "werwer",
		"r":   "werwer",
	}
	jsonResponse(res, comm)
}

// func (s *Server) Assets(res http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(res, req, req.URL.Path[1:])
// }

// Todo CRUD

func (s *Server) TodoIndex(res http.ResponseWriter, req *http.Request) {
	var todos []*Todo

	rows, err := s.Db.Query("SELECT id, name, description, state FROM todos")
	errorCheck(res, err)
	for rows.Next() {
		todo := &Todo{}
		rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.State)
		todos = append(todos, todo)
	}
	rows.Close()

	jsonResponse(res, todos)
}

func (s *Server) TodoCreate(res http.ResponseWriter, req *http.Request) {
	todo := &Todo{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		fmt.Println("ERROR decoding JSON - ", err)
		return
	}
	defer req.Body.Close()

	result, err := s.Db.Exec("INSERT INTO todos(name, description, state) VALUES(?, ?, ?)", todo.Name, todo.Description, todo.State)
	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	Id64, err := result.LastInsertId()
	Id := int(Id64)
	todo = &Todo{Id: Id}

	s.Db.QueryRow("SELECT state, name, description FROM todos WHERE Id=?", todo.Id).Scan(&todo.State, &todo.Name, &todo.Description)

	jsonResponse(res, todo)
}

func (s *Server) TodoShow(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Render todo json")
}

func (s *Server) TodoUpdate(res http.ResponseWriter, req *http.Request) {
	todoParams := &Todo{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&todoParams)
	if err != nil {
		fmt.Println("ERROR decoding JSON - ", err)
		return
	}

	_, err = s.Db.Exec("UPDATE todos SET name = ?, description = ?, state = ? WHERE id = ?", todoParams.Name, todoParams.Description, todoParams.State, todoParams.Id)

	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	todo := &Todo{Id: todoParams.Id}
	err = s.Db.QueryRow("SELECT state, name, description FROM todos WHERE id=?", todo.Id).Scan(&todo.State, &todo.Name, &todo.Description)
	if err != nil {
		fmt.Println("ERROR reading from db - ", err)
	}

	jsonResponse(res, todo)
}

func (s *Server) TodoDelete(res http.ResponseWriter, req *http.Request) {
	r, _ := regexp.Compile(`\d+$`)
	id := r.FindString(req.URL.Path)
	s.Db.Exec("DELETE FROM todos WHERE id=?", id)
	res.WriteHeader(200)
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

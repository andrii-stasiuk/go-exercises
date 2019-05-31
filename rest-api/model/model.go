package model

import (
	"database/sql"
	"fmt"
)

// store "context" values and connections in the server struct
type Model struct {
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

func (s *Model) Index() []*Todo {
	var todos []*Todo

	rows, _ /*err*/ := s.Db.Query("SELECT id, name, description, state FROM todos")
	//errors.ErrorCheck(w, err)
	for rows.Next() {
		todo := &Todo{}
		rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.State)
		todos = append(todos, todo)
	}
	rows.Close()
	return todos
}

func (s *Model) Delete(id uint64) error {
	_, err := s.Db.Exec("DELETE FROM todos WHERE id=?", id)
	return err
}

func (s *Model) Show(id uint64) Todo {

	todo := &Todo{}
	err := s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id).Scan(&todo.Id, &todo.State, &todo.Name, &todo.Description, &todo.Created_at, &todo.Updated_at)
	if err != nil {
		fmt.Println("ERROR reading from db - ", err)
	}

	return *todo
}

func (s *Model) Create(name, descr string) Todo {
	todo := &Todo{}
	todo.Name = name
	todo.Description = descr

	result, err := s.Db.Exec("INSERT INTO todos(name, description) VALUES(?, ?)", todo.Name, todo.Description)
	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	id64, err := result.LastInsertId()
	//print(id64)
	s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id64).Scan(&todo.Id, &todo.State, &todo.Name, &todo.Description, &todo.Created_at, &todo.Updated_at)
	return *todo
}

func (s *Model) Update(id uint64, name, descr string) Todo {
	todo := &Todo{}
	/*result*/ _, err := s.Db.Exec("UPDATE todos SET name = ?, description = ? WHERE id = ?", name, descr, id)
	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}
	//id64, err := result.RowsAffected()
	//print(id64)
	err = s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id).Scan(&todo.Id, &todo.State, &todo.Name, &todo.Description, &todo.Created_at, &todo.Updated_at)
	if err != nil {
		fmt.Println("ERROR reading from db - ", err)
	}
	// make if name = todo.Name, ...
	return *todo
}

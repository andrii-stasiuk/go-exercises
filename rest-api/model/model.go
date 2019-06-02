package model

import (
	"database/sql"
	"time"
)

// type Status struct {
// 	Id     uint64
// 	Status string
// }

// Statuses - a map to store the statuses with the ID as the key
// var Statuses = map[uint64]string{
// 	1: "created",
// 	2: "in process",
// 	3: "resolved",
// }

// Model store "context" values and connections in the server struct
type Model struct {
	Db *sql.DB
}

// Todo main identifier
type Todo struct {
	ID          int    `json:"id,sting"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Index method
func (s *Model) Index() ([]*Todo, error) {
	var todos []*Todo
	rows, err := s.Db.Query("SELECT id, name, description, state, created_at, updated_at FROM todos")
	if err != nil {
		return []*Todo{}, err
	}
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.State, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return []*Todo{}, err
		}
		todos = append(todos, todo)
	}
	err = rows.Close()
	if err != nil {
		return []*Todo{}, err
	}
	return todos, nil
}

// Show method
func (s *Model) Show(id string) (Todo, error) {
	todo := Todo{}
	row := s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id)
	err := row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	return todo, err
}

// Delete method
func (s *Model) Delete(id string) (bool, error) {
	result, err := s.Db.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return false, err
	}
	id64, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if id64 > 0 {
		return true, nil
	}
	return false, err
}

// Create method
func (s *Model) Create(name, descr string, state string) (Todo, error) {
	todo := Todo{}
	result, err := s.Db.Exec("INSERT INTO todos(name, description, state, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", name, descr, state, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return Todo{}, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	row := s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id64)
	err = row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// Update method
func (s *Model) Update(id string, name, descr string, state string) (Todo, error) {
	todo := Todo{}
	_, err := s.Db.Exec("UPDATE todos SET name = ?, description = ?, state = ?, updated_at = ? WHERE id = ?", name, descr, state, time.Now().Unix(), id)
	if err != nil {
		return Todo{}, err
	}
	// id64, err := result.RowsAffected()
	// if err != nil {
	// 	return Todo{}, err
	// }
	row := s.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id)
	err = row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

/*Package model Todo*/
package model

import (
	"database/sql"
	"time"
)

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

//States - a map to store the states of Todo with the ID as the key
var States = map[string]string{
	"1": "created",
	"2": "wait",
	"3": "canceled",
	"4": "blocked",
	"5": "in process/doing",
	"6": "review",
	"7": "done",
	"8": "archived",
}

// Index method
func (m *Model) Index() ([]*Todo, error) {
	var todos []*Todo
	rows, err := m.Db.Query("SELECT id, name, description, state, created_at, updated_at FROM todos")
	if err != nil {
		return []*Todo{}, err
	}
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.State, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return []*Todo{}, err
		}
		todo.State = States[todo.State] // Shows the State in human-readable form
		todos = append(todos, todo)
	}
	err = rows.Close()
	if err != nil {
		return []*Todo{}, err
	}
	return todos, nil
}

// Show method
func (m *Model) Show(id string) (*Todo, error) {
	var todo Todo
	row := m.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id)
	err := row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Shows the State in human-readable form
	return &todo, err
}

// Delete method
func (m *Model) Delete(id string) (bool, error) {
	result, err := m.Db.Exec("DELETE FROM todos WHERE id=?", id)
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
func (m *Model) Create(todo *Todo) (*Todo, error) {
	result, err := m.Db.Exec("INSERT INTO todos(name, description, state, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", todo.Name, todo.Description, todo.State, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return &Todo{}, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return &Todo{}, err
	}
	row := m.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id64)
	err = row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Shows the State in human-readable form
	return todo, nil
}

// Update method
func (m *Model) Update(id string, todo *Todo) (*Todo, error) {
	//todo := Todo{}
	_, err := m.Db.Exec("UPDATE todos SET name = ?, description = ?, state = ?, updated_at = ? WHERE id = ?", todo.Name, todo.Description, todo.State, time.Now().Unix(), id)
	if err != nil {
		return &Todo{}, err
	}
	// id64, err := result.RowsAffected()
	// if err != nil {
	// 	return Todo{}, err
	// }
	row := m.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=?", id)
	err = row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Shows the State in human-readable form
	return todo, nil
}

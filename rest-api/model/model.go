/*Package model Todo*/
package model

import (
	"database/sql"
)

// Todos model store "context" values and connections in the server struct
type Todos struct {
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

// States - a map to store the states of Todo with the ID as the key,
// it can be stored in other related table "states" in the future
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

// New gets the address of the database as parameter and returns new Model struct
func New(db *sql.DB) Todos {
	return Todos{Db: db}
}

// Index method to get all the records in a table
func (m Todos) Index() ([]*Todo, error) {
	var todos []*Todo
	rows, err := m.Db.Query("SELECT id, name, description, state, created_at, updated_at FROM todos ORDER BY id")
	if err != nil {
		return []*Todo{}, err
	}
	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.State, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return []*Todo{}, err
		}
		todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
		todos = append(todos, todo)
	}
	err = rows.Close()
	return todos, err
}

// Show method to get a specific record from a table
func (m Todos) Show(id string) (*Todo, error) {
	var todo Todo
	row := m.Db.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=$1", id)
	err := row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	return &todo, err
}

// Delete method to delete a specific record from a table
func (m Todos) Delete(id string) (*Todo, error) {
	var todo Todo
	sqlStatement := `DELETE FROM todos WHERE id=$1 RETURNING id, state, name, description, created_at, updated_at`
	err := m.Db.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	return &todo, err
}

// Create method to create a record in the table
func (m Todos) Create(todo *Todo) (*Todo, error) {
	sqlStatement := "INSERT INTO todos (name, description, state) VALUES($1, $2, $3) RETURNING id, created_at, updated_at"
	err := m.Db.QueryRow(sqlStatement, todo.Name, todo.Description, todo.State).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	return todo, err
}

// Update method to change the record in the table
func (m Todos) Update(todo *Todo) (*Todo, error) {
	sqlStatement := "UPDATE todos SET name = $1, description = $2, state = $3, updated_at = now() WHERE id=$4 RETURNING id, created_at, updated_at"
	err := m.Db.QueryRow(sqlStatement, todo.Name, todo.Description, todo.State, todo.ID).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return &Todo{}, err
	}
	todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	return todo, err
}

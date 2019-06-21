/*Package model Todo*/
package todomodel

import (
	"strconv"

	"github.com/jmoiron/sqlx"
)

// Todos model store "context" values and connections in the server struct
type Todos struct {
	DB *sqlx.DB
}

// Todo main identifier
type Todo struct {
	ID          uint64 `json:"id,sting"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
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

// NewTodo gets the address of the database as parameter and returns new Model struct
func NewTodo(db *sqlx.DB) Todos {
	return Todos{DB: db}
}

// Index method to get all the records in a table
func (t Todos) Index() ([]Todo, error) {
	todos := []Todo{}
	sqlStatement := `SELECT todos.id, todos.name, todos.description, states.state, todos.created_at, todos.updated_at 
	FROM todos LEFT JOIN states ON todos.state=states.id ORDER BY id`
	err := t.DB.Select(&todos, sqlStatement)
	return todos, err
}

// Show method to get a specific record from a table
func (t Todos) Show(id string) (Todo, error) {
	todo := Todo{}
	sqlStatement := `SELECT todos.id, todos.name, todos.description, states.state, todos.created_at, todos.updated_at
	FROM todos LEFT JOIN states ON todos.state=states.id WHERE todos.id=$1`
	err := t.DB.Get(&todo, sqlStatement, id)
	return todo, err
}

// Delete method to delete a specific record from a table
func (t Todos) Delete(id string) (Todo, error) {
	todo := Todo{}
	sqlStatement := "DELETE FROM todos WHERE id=$1 RETURNING id, state, name, description, created_at, updated_at"
	err := t.DB.Get(&todo, sqlStatement, id)
	return todo, err
}

// Create method to create a record in the table
func (t Todos) Create(todo Todo) (Todo, error) {
	sqlStatement := "INSERT INTO todos (name, description, state) VALUES($1, $2, $3) RETURNING id, created_at, updated_at"
	err := t.DB.Get(&todo, sqlStatement, todo.Name, todo.Description, todo.State)
	return todo, err
}

// Update method to change the record in the table
func (t Todos) Update(todo Todo) (Todo, error) {
	sqlStatement := "UPDATE todos SET "
	if todo.Name != "" {
		sqlStatement += "name = '" + todo.Name + "', "
	}
	if todo.Description != "" {
		sqlStatement += "description = '" + todo.Description + "', "
	}
	if todo.State != "" {
		sqlStatement += "state = " + todo.State + ", "
	}
	sqlStatement += "updated_at = now() WHERE id=" + strconv.FormatUint(todo.ID, 10) + " RETURNING name, description, state, created_at, updated_at"
	err := t.DB.Get(&todo, sqlStatement)
	return todo, err
}

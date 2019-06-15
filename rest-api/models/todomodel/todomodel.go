/*Package model Todo*/
package todomodel

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Todos model store "context" values and connections in the server struct
type Todos struct {
	DB *gorm.DB
}

// Todo main identifier
type Todo struct {
	ID          uint64    `json:"id,sting" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	State       string    `json:"state" gorm:"state"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
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
func NewTodo(db *gorm.DB) Todos {
	return Todos{DB: db}
}

// Index method to get all the records in a table
func (t Todos) Index() ([]Todo, error) {
	// var todos []Todo
	// rows, err := m.DB.Query("SELECT id, name, description, state, created_at, updated_at FROM todos ORDER BY id")
	// if err != nil {
	// 	return []Todo{}, err
	// }
	// for rows.Next() {
	// 	todo := Todo{}
	// 	err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.State, &todo.CreatedAt, &todo.UpdatedAt)
	// 	if err != nil {
	// 		return []Todo{}, err
	// 	}
	// 	todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	// 	todos = append(todos, todo)
	// }
	// err = rows.Close()
	// return todos, err
	var todos []Todo
	return todos, t.DB.Order("id").Find(&todos).Error
}

// Show method to get a specific record from a table
func (t Todos) Show(id string) (Todo, error) {
	// var todo Todo
	// row := m.DB.QueryRow("SELECT id, state, name, description, created_at, updated_at FROM todos WHERE id=$1", id)
	// err := row.Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	// if err != nil {
	// 	return Todo{}, err
	// }
	// todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	// return todo, err
	var todo Todo
	return todo, t.DB.Where("id = ?", id).First(&todo).Error
}

// Delete method to delete a specific record from a table
func (t Todos) Delete(id string) (Todo, error) {
	// var todo Todo
	// sqlStatement := `DELETE FROM todos WHERE id=$1 RETURNING id, state, name, description, created_at, updated_at`
	// err := m.DB.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.State, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	// if err != nil {
	// 	return Todo{}, err
	// }
	// todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	// return todo, err
	var todo Todo
	err := t.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return Todo{}, err
	}
	return todo, t.DB.Where("id = ?", id).Delete(Todo{}).Error // don't returns error when field doesn't exist
}

// Create method to create a record in the table
func (t Todos) Create(todo Todo) (Todo, error) {
	// sqlStatement := "INSERT INTO todos (name, description, state) VALUES($1, $2, $3) RETURNING id, created_at, updated_at"
	// err := m.DB.QueryRow(sqlStatement, todo.Name, todo.Description, todo.State).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	// if err != nil {
	// 	return Todo{}, err
	// }
	// todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	// return todo, err
	return todo, t.DB.Create(&todo).Error
}

// Update method to change the record in the table
func (t Todos) Update(todo Todo) (Todo, error) {
	// sqlStatement := "UPDATE todos SET name = $1, description = $2, state = $3, updated_at = now() WHERE id=$4 RETURNING id, created_at, updated_at"
	// err := m.DB.QueryRow(sqlStatement, todo.Name, todo.Description, todo.State, todo.ID).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
	// if err != nil {
	// 	return Todo{}, err
	// }
	// todo.State = States[todo.State] // Changes the State to show it by API in human-readable form (reserved for future purposes)
	// return todo, err
	err := t.DB.Where("id = ?", todo.ID).First(&Todo{}).Error // Check if exists
	if err != nil {
		return Todo{}, err
	}
	err = t.DB.Where("id = ?", todo.ID).Save(&todo).Error // Save() doesn't return "created_at" field
	if err != nil {
		return Todo{}, err
	}
	return todo, t.DB.First(&todo).Error // First() returns all the fields
}

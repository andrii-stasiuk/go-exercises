/*Package handlers Todo*/
package todo

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/todomodel"
)

type todoInterface interface {
	Index(userID string) ([]todomodel.Todo, error)
	Show(todoID string, userID string) (todomodel.Todo, error)
	Delete(todoID string, userID string) (todomodel.Todo, error)
	Create(todo todomodel.Todo) (todomodel.Todo, error)
	Update(todo todomodel.Todo) (todomodel.Todo, error)
}

// Handlers structure for handling requests
type TodoHandlers struct {
	SQL todoInterface
}

// New is a constructor of "Handlers" that gets "Model" type Model as an argument and returns "Handlers" type Handlers
func New(ti todoInterface) TodoHandlers {
	return TodoHandlers{SQL: ti}
}

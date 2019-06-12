/*Package handlers Todo*/
package todo

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/models/todomodel"
)

type todoInterface interface {
	Index() ([]*todomodel.Todo, error)
	Show(id string) (*todomodel.Todo, error)
	Delete(id string) (*todomodel.Todo, error)
	Create(todo *todomodel.Todo) (*todomodel.Todo, error)
	Update(todo *todomodel.Todo) (*todomodel.Todo, error)
}

// Handlers structure for handling requests
type TodoHandlers struct {
	SQL todoInterface
}

// New is a constructor of "Handlers" that gets "Model" type Model as an argument and returns "Handlers" type Handlers
func New(ti todoInterface) TodoHandlers {
	return TodoHandlers{SQL: ti}
}

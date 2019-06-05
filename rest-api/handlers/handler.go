/*Package handler Todo*/
package handlers

import (
	"github.com/andrii-stasiuk/go-exercises/rest-api/model"
)

type modelInterface interface {
	Index() ([]*model.Todo, error)
	Show(id string) (*model.Todo, error)
	Delete(id string) (*model.Todo, error)
	Create(todo *model.Todo) (*model.Todo, error)
	Update(id string, todo *model.Todo) (*model.Todo, error)
}

// Handlers structure for handling requests
type Handlers struct {
	SQL modelInterface
}

// New is a constructor of "Handlers", it gets "Model" type Model as an argument and returns "Handlers" type Handlers
func New(mi modelInterface) Handlers {
	return Handlers{SQL: mi}
}

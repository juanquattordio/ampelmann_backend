package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Cliente interface {
	Save(cliente entities.Cliente) error
	GetLastID() (int64, error)
	//Search(context context.Context, command search_refinances.Request) ([]entities.Refinance, *entities.Page, error)
}

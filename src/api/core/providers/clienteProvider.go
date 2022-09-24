package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Cliente interface {
	Save(cliente entities.Cliente) error
	GetLastID() (int64, error)
	Search(id *int64, cuit *string) (*entities.Cliente, error)
	Update(cliente *entities.Cliente) error
}

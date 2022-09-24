package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Deposito interface {
	Save(deposito entities.Deposito) error
	GetLastID() (int64, error)
	Search(id *int64, descripcion *string) (*entities.Deposito, error)
	Update(deposito *entities.Deposito) error
}

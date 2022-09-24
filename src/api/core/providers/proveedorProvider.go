package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Proveedor interface {
	Save(proveedor entities.Proveedor) error
	GetLastID() (int64, error)
	Search(id *int64, cuit *string) (*entities.Proveedor, error)
	Update(proveedor *entities.Proveedor) error
}

package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Insumo interface {
	Save(cliente entities.Insumo) error
	GetLastID() (int64, error)
	Search(id *int64, nombre *string) (*entities.Insumo, error)
}

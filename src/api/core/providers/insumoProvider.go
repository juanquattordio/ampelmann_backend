package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Insumo interface {
	Save(insumo entities.Insumo) error
	GetLastID() (int64, error)
	Search(id *int64, nombre *string) (*entities.Insumo, error)
	Update(insumo *entities.Insumo) error
}

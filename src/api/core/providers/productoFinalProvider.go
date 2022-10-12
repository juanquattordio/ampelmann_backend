package providers

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ProductoFinal interface {
	Save(productoFinal entities.ProductoFinal) (int64, error)
	Search(id *int64, descripcion *string) (*entities.ProductoFinal, error)
	Update(productoFinal *entities.ProductoFinal) error
}

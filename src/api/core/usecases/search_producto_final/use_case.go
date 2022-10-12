package search_producto_final

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id *int64, descripcion *string) (*entities.ProductoFinal, error)
}

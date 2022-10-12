package update_producto_final

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id int64, request update_producto_final.RequestUpdate) (*entities.ProductoFinal, error)
}

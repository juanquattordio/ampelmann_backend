package create_producto_final

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_producto_final.Request) (*entities.ProductoFinal, error)
}

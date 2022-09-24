package update_proveedor

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id int64, request update_proveedor.Request) (*entities.Proveedor, error)
}

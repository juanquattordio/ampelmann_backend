package create_proveedor

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_proveedor.Request) (*entities.Proveedor, error)
}

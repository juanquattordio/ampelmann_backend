package create_insumo

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_insumo.Request) (*entities.Insumo, error)
}

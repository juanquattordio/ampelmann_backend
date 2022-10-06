package movimiento_depositos

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request movimiento_depositos.Request) (*entities.MovimientoHeader, error)
}

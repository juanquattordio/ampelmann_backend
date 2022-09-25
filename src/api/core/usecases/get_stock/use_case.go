package get_stock

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, idInsumo *int64, idDeposito *int64) (*entities.Insumo, *entities.Deposito, float64, error)
}

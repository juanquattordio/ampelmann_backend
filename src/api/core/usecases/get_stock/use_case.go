package get_stock

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	GetStockByInsumo(context context.Context, idInsumo *int64, idDeposito *int64) (*entities.Insumo, *entities.Deposito, error)
	GetStockByDeposito(context context.Context, idDeposito *int64) (*entities.Deposito, []entities.Insumo, error)
}

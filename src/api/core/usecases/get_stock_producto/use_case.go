package get_stock_producto

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	GetStockByProducto(context context.Context, idProducto *int64, idDeposito *int64) (*entities.ProductoFinal, *entities.Deposito, error)
	GetStockByDeposito(context context.Context, idDeposito *int64) (*entities.Deposito, []entities.ProductoFinal, error)
}

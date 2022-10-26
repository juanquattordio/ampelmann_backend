package providers

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type StockProducto interface {
	GetStockProducto(idProducto *int64, idDeposito *int64) (float64, error)
	GetStockDeposito(ctx context.Context, idDeposito *int64) ([]entities.ProductoFinal, error)
	//MovimientoDepositos(ctx context.Context, movimiento *entities.MovimientoHeader) error
	//UpdateStock(tx *sqlx.Tx, idProducto *int64, idDeposito *int64, cantidad float64) error
}

package providers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Stock interface {
	GetStockInsumo(idInsumo *int64, idDeposito *int64) (float64, error)
	GetStockDeposito(ctx context.Context, idDeposito *int64) ([]entities.Insumo, error)
	MovimientoDepositos(ctx context.Context, movimiento *entities.MovimientoHeader, tipoArticulo string) error
	UpdateStock(tx *sqlx.Tx, idInsumo *int64, idDeposito *int64, cantidad float64, tipoArticulo string) error
}

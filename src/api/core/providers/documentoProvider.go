package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Documento interface {
	CreateHeaderMovimientoDepositos(tx *sqlx.Tx, movimiento *entities.MovimientoHeader) (int64, error)
	CreateLineMovimientoDepositos(tx *sqlx.Tx, idHeader int64, idLinea int, idInsumo *int64, cantidad *float64, observaciones *string) error
	CreateFacturaCompra(factura *entities.FacturaCompraHeader) error
	CreateHeaderFacturaCompra(tx *sqlx.Tx, factura *entities.FacturaCompraHeader) (int64, error)
	CreateLineFacturaCompra(tx *sqlx.Tx, idHeader int64, idLinea int, idInsumo *int64, cantidad *float64, precioUnitario *float64, observaciones *string) error
}

package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type DocumentoProducto interface {
	CreateHeaderMovimientoDepositos(tx *sqlx.Tx, movimiento *entities.MovimientoProductoHeader) (int64, error)
	CreateLineMovimientoDepositos(tx *sqlx.Tx, idHeader int64, idLinea int, idArticulo *int64, cantidad *float64, observaciones *string) error
	CreateFacturaVenta(factura *entities.FacturaVentaHeader) error
	CreateHeaderFacturaCompra(tx *sqlx.Tx, factura *entities.FacturaVentaHeader) (int64, error)
	CreateLineFacturaCompra(tx *sqlx.Tx, idHeader int64, idLinea int, idArticulo *int64, cantidad *float64, precioUnitario *float64, observaciones *string) error
}

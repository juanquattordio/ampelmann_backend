package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Documento interface {
	CreateHeaderMovimientoDepositos(tx *sqlx.Tx, movimiento *entities.MovimientoHeader) (int64, error)
	CreateLineMovimientoDepositos(tx *sqlx.Tx, idHeader int64, idLinea int, idInsumo *int64, cantidad *float64, observaciones *string) error
}

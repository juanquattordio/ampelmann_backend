package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Documento interface {
	CreateMovimientoDepositos(tx *sqlx.Tx, movimiento *entities.MovimientoHeader) error
}

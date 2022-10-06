package documento

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Documento {
	repo := Repository{
		db: db,
	}
	return &repo
}

func (r *Repository) CreateMovimientoDepositos(tx *sqlx.Tx, movimiento *entities.MovimientoHeader) error {
	var err error
	if tx == nil {
		//_, err = r.db.Query(updateStock, &idDeposito, &idInsumo, &cantidad)
	} else {
		_, err = tx.Query(insertMovInsumoHeader, &movimiento.IdDepositoOrigen, &movimiento.IdDepositoDestino, &movimiento.Fecha)
		if err != nil {
			tx.Rollback()
			return errors.NewInternalServer("Fallo al crear documento")
		}
	}
	return err
}

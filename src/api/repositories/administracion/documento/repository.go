package documento

import (
	"fmt"
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

func (r *Repository) CreateHeaderMovimientoDepositos(tx *sqlx.Tx, movimiento *entities.MovimientoHeader) (int64, error) {
	var idHeader int64
	if tx == nil {
		stmt, err := r.db.Prepare(insertMovInsumoHeader)
		result, err := stmt.Exec(&movimiento.IdDepositoOrigen, &movimiento.IdDepositoDestino, &movimiento.Fecha)
		if err != nil {
			return 0, errors.NewInternalServer("Fallo al crear documento")
		}
		idHeader, err = result.LastInsertId()
	} else {
		stmt, err := tx.Prepare(insertMovInsumoHeader)
		result, err := stmt.Exec(&movimiento.IdDepositoOrigen, &movimiento.IdDepositoDestino, &movimiento.Fecha)
		if err != nil {
			tx.Rollback()
			return 0, errors.NewInternalServer("Fallo al crear documento")
		}
		idHeader, err = result.LastInsertId()
	}
	return idHeader, nil
}

func (r *Repository) CreateLineMovimientoDepositos(tx *sqlx.Tx, idHeader int64, idLinea int, idInsumo *int64, cantidad *float64,
	obseraciones *string) error {
	idLinea += 1
	if tx == nil {
		stmt, err := r.db.Prepare(insertMovInsumoLine)
		_, err = stmt.Exec(idHeader, idLinea, &idInsumo, &cantidad, &obseraciones)
		if err != nil {
			return errors.NewInternalServer(fmt.Sprintf("Fallo al crear linea %d de header id %d", idLinea, idHeader))
		}
	} else {
		stmt, err := tx.Prepare(insertMovInsumoLine)
		_, err = stmt.Exec(idHeader, idLinea, &idInsumo, &cantidad, &obseraciones)
		if err != nil {
			tx.Rollback()
			return errors.NewInternalServer(fmt.Sprintf("Fallo al crear documento %d de header id %d", idLinea, idHeader))
		}
	}
	return nil
}

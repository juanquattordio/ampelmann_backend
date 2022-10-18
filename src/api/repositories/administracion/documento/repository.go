package documento

import (
	"database/sql"
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
		result, err := stmt.Exec(&movimiento.IdDepositoOrigen, &movimiento.IdDepositoDestino, &movimiento.Fecha, movimiento.CausaMovimiento)
		if err != nil {
			return 0, errors.NewInternalServer("Fallo al crear documento")
		}
		idHeader, err = result.LastInsertId()
	} else {
		stmt, err := tx.Prepare(insertMovInsumoHeader)
		result, err := stmt.Exec(&movimiento.IdDepositoOrigen, &movimiento.IdDepositoDestino, &movimiento.Fecha, movimiento.CausaMovimiento)
		if err != nil {
			return 0, errors.NewInternalServer("Fallo al crear documento")
		}
		idHeader, err = result.LastInsertId()
	}
	return idHeader, nil
}

func (r *Repository) CreateLineMovimientoDepositos(tx *sqlx.Tx, idHeader int64, idLinea int, idInsumo *int64, cantidad *float64,
	observaciones *string) error {
	idLinea += 1
	if tx == nil {
		stmt, err := r.db.Prepare(insertMovInsumoLine)
		_, err = stmt.Exec(idHeader, idLinea, &idInsumo, &cantidad, &observaciones)
		if err != nil {
			return errors.NewInternalServer(fmt.Sprintf("Fallo al crear linea %d de header id %d", idLinea, idHeader))
		}
	} else {
		stmt, err := tx.Prepare(insertMovInsumoLine)
		_, err = stmt.Exec(idHeader, idLinea, &idInsumo, &cantidad, &observaciones)
		if err != nil {
			return errors.NewInternalServer(fmt.Sprintf("Fallo al crear documento %d de header id %d", idLinea, idHeader))
		}
	}
	return nil
}

func (r *Repository) CreateFacturaCompra(factura *entities.FacturaCompraHeader) error {
	tx := r.db.MustBegin()
	var err error
	var idHeader int64

	// inserta header en tabla
	if idHeader, err = r.CreateHeaderFacturaCompra(tx, factura); err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			return errors.NewInternalServer("Fallo en el rollback de la transacción")
		}
		return errors.NewInternalServer("Fallo en crear header factura de compra. Se hace rollback")
	}
	factura.IdFactura = idHeader

	// inserta las líneas en la tabla
	for i, linea := range factura.Lineas {
		if err = r.CreateLineFacturaCompra(tx, idHeader, i, &linea.IdArticulo, &linea.Cantidad, &linea.PrecioUnitario,
			&linea.Observaciones); err != nil {
			break
		}
		factura.Lineas[i].IdLinea = int64(i + 1)
	}
	if err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			return errors.NewInternalServer("Fallo en el rollback de la transacción")
		}
		return errors.NewInternalServer("Fallo en Update de Stock. Se hace rollback")
	}

	if err = tx.Commit(); err != nil {
		return errors.NewInternalServer("Fallo en el commit de la transacción")
	}
	return nil
}

func (r *Repository) CreateHeaderFacturaCompra(tx *sqlx.Tx, factura *entities.FacturaCompraHeader) (int64, error) {
	var (
		idHeader int64
		stmt     *sql.Stmt
		result   sql.Result
		err      error
	)
	if tx == nil {
		stmt, err = r.db.Prepare(insertFacturaCompraHeader)
	} else {
		stmt, err = tx.Prepare(insertFacturaCompraHeader)
	}
	if factura.FechaOrigen.IsZero() {
		result, err = stmt.Exec(&factura.IdProveedor, &factura.IdFacturaProveedor, nil, &factura.Fecha, &factura.ImporteTotal)
	} else {
		result, err = stmt.Exec(&factura.IdProveedor, &factura.IdFacturaProveedor, &factura.FechaOrigen, &factura.Fecha, &factura.ImporteTotal)
	}
	if err != nil {
		return 0, errors.NewInternalServer("Fallo al crear documento")
	}
	idHeader, err = result.LastInsertId()
	return idHeader, nil
}

func (r *Repository) CreateLineFacturaCompra(tx *sqlx.Tx, idHeader int64, idLinea int, idInsumo *int64,
	cantidad *float64, precioUnitario *float64, observaciones *string) error {
	var (
		stmt *sql.Stmt
		err  error
	)
	idLinea += 1

	if tx == nil {
		stmt, err = r.db.Prepare(insertFacturaCompraLine)
	} else {
		stmt, err = tx.Prepare(insertFacturaCompraLine)
	}
	_, err = stmt.Exec(idHeader, idLinea, &idInsumo, &cantidad, &precioUnitario, &observaciones)
	if err != nil {
		return errors.NewInternalServer(fmt.Sprintf("Fallo al crear documento %d de header id %d", idLinea, idHeader))
	}

	return nil
}

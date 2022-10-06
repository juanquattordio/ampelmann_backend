package stock

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db                *sqlx.DB
	documentoProvider providers.Documento
}

func NewRepository(db *sqlx.DB, documentoProvider providers.Documento) providers.Stock {
	repo := Repository{
		db:                db,
		documentoProvider: documentoProvider,
	}
	return &repo
}

func (r *Repository) GetStockInsumo(idInsumo *int64, idDeposito *int64) (float64, error) {
	if idDeposito == nil {
		dbStock := new(stockInsumo)
		rows, err := r.db.Query(sumStockByInsumo, idInsumo)
		if err != nil {
			return 0, err
		}
		for rows.Next() {
			_ = rows.Scan(&dbStock.IdInsumo, &dbStock.Stock)
		}
		return dbStock.Stock, nil
	}

	// caso de stockInsumo por insumo y deposito
	whereConditions, args := buildSearchWhere(idInsumo, idDeposito)
	stockScript := getStockInsumoDeposito + whereConditions
	dbStock := new(stockInsumo)

	err := r.db.Get(dbStock, stockScript, args...)

	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return dbStock.Stock, nil

}

func (r *Repository) GetStockDeposito(ctx context.Context, idDeposito *int64) ([]entities.Insumo, error) {
	var dbStockDeposito []stockDeposito
	err := r.db.SelectContext(ctx, &dbStockDeposito, getStockByDeposito, idDeposito)
	if err != nil && !goErrors.Is(err, sql.ErrNoRows) || dbStockDeposito == nil {
		return nil, errors.NewInternalServer("Deposito sin stock o inexistente")
	}
	var insumos []entities.Insumo
	for _, insumoDB := range dbStockDeposito {
		//_ = rows.Scan(&insumoDB.IdDeposito, &insumoDB.IdInsumo, &insumoDB.NombreInsumo, &insumoDB.Stock)
		insumo := insumoDB.toEntity()
		insumos = append(insumos, insumo)
	}

	return insumos, nil
}

func (r *Repository) MovimientoDepositos(ctx context.Context, header *entities.MovimientoHeader) error {
	tx := r.db.MustBegin()
	var err error

	// inserta header en tabla
	if err = r.documentoProvider.CreateMovimientoDepositos(tx, header); err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			return errors.NewInternalServer("Fallo en el rollback de la transacción")
		}
		return errors.NewInternalServer("Fallo en crear header movimiento insumos. Se hace rollback")
	}

	for _, linea := range header.Lineas {
		if err = r.UpdateStock(tx, &linea.IdInsumo, &header.IdDepositoOrigen, -(linea.Cantidad)); err != nil {
			break
		}
		if err = r.UpdateStock(tx, &linea.IdInsumo, &header.IdDepositoDestino, linea.Cantidad); err != nil {
			break
		}
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

func (r *Repository) UpdateStock(tx *sqlx.Tx, idInsumo *int64, idDeposito *int64, cantidad float64) error {
	var err error
	if tx == nil {
		_, err = r.db.Query(updateStock, &idDeposito, &idInsumo, &cantidad)
	} else {
		_, err = tx.Query(updateStock, &idDeposito, &idInsumo, &cantidad)
		if err != nil {
			tx.Rollback()
			return errors.NewInternalServer("Fallo al actualizar stock")
		}
	}

	return err
}

func buildSearchWhere(idInsumo *int64, idDeposito *int64) (query string, args []interface{}) {
	if idInsumo != nil {
		query += " AND id_insumo = ?"
		args = append(args, idInsumo)
	}
	if idDeposito != nil {
		query += " AND id_deposito = ?"
		args = append(args, idDeposito)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}

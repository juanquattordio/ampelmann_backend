package stock_producto

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

// func NewRepository(db *sqlx.DB, documentoProvider providers.DocumentoProducto) providers.StockProducto {
func NewRepository(db *sqlx.DB) providers.StockProducto {
	repo := Repository{
		db: db,
		//documentoProvider: documentoProvider,
	}
	return &repo
}

func (r *Repository) GetStockProducto(idProducto *int64, idDeposito *int64) (float64, error) {
	if idDeposito == nil {
		dbStock := new(stockProducto)
		rows, err := r.db.Query(sumStockByProducto, idProducto)
		if err != nil {
			return 0, err
		}
		for rows.Next() {
			_ = rows.Scan(&dbStock.IdProducto, &dbStock.Stock)
		}
		return dbStock.Stock, nil
	}

	// caso de stockProducto por producto y deposito
	whereConditions, args := buildSearchWhere(idProducto, idDeposito)
	stockScript := getStockProductoDeposito + whereConditions
	dbStock := new(stockProducto)

	err := r.db.Get(dbStock, stockScript, args...)

	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return dbStock.Stock, nil

}

func (r *Repository) GetStockDeposito(ctx context.Context, idDeposito *int64) ([]entities.ProductoFinal, error) {
	var dbStockDeposito []stockDeposito
	err := r.db.SelectContext(ctx, &dbStockDeposito, getStockByDeposito, idDeposito)
	if err != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, errors.NewInternalServer("Deposito inexistente")
	}
	if dbStockDeposito == nil {
		return nil, nil
	}
	var productos []entities.ProductoFinal
	for _, productoDB := range dbStockDeposito {
		producto := productoDB.toEntity()
		productos = append(productos, producto)
	}

	return productos, nil
}

//func (r *Repository) MovimientoDepositos(ctx context.Context, header *entities.MovimientoProductoHeader) error {
//	tx := r.db.MustBegin()
//	var err error
//	var idHeader int64
//
//	// inserta header en tabla
//	if idHeader, err = r.documentoProvider.CreateHeaderMovimientoDepositos(tx, header); err != nil {
//		if errRollBack := tx.Rollback(); errRollBack != nil {
//			return errors.NewInternalServer("Fallo en el rollback de la transacción")
//		}
//		return errors.NewInternalServer("Fallo en crear header movimiento productos. Se hace rollback")
//	}
//	header.IdHeader = idHeader
//
//	// por cada linea del comprobante se actualiza stocks e inserta linea de comprobante en tabla
//	for i, linea := range header.Lineas {
//		if err = r.UpdateStock(tx, &linea.IdProducto, &header.IdDepositoOrigen, -(linea.Cantidad)); err != nil {
//			break
//		}
//		if err = r.UpdateStock(tx, &linea.IdProducto, &header.IdDepositoDestino, linea.Cantidad); err != nil {
//			break
//		}
//		if err = r.documentoProvider.CreateLineMovimientoDepositos(tx, idHeader, i, nil, &linea.IdProducto, &linea.Cantidad,
//			&linea.Observaciones); err != nil {
//			break
//		}
//		header.Lineas[i].IdLinea = int64(i + 1)
//	}
//	if err != nil {
//		if errRollBack := tx.Rollback(); errRollBack != nil {
//			return errors.NewInternalServer("Fallo en el rollback de la transacción")
//		}
//		return errors.NewInternalServer("Fallo en Update de Stock. Se hace rollback")
//	}
//
//	if err = tx.Commit(); err != nil {
//		return errors.NewInternalServer("Fallo en el commit de la transacción")
//	}
//	return nil
//}

func (r *Repository) UpdateStock(tx *sqlx.Tx, idProducto *int64, idDeposito *int64, cantidad float64) error {
	var err error
	if tx == nil {
		_, err = r.db.NamedExec(updateStock, toDBEntity(*idDeposito, *idProducto, cantidad))
		if err != nil {
			return err
		}
	} else {
		_, err = tx.NamedExec(updateStock, toDBEntity(*idDeposito, *idProducto, cantidad))
		if err != nil {
			tx.Rollback()
			return errors.NewInternalServer("Fallo al actualizar stock")
		}
	}

	return err
}

func buildSearchWhere(idProducto *int64, idDeposito *int64) (query string, args []interface{}) {
	if idProducto != nil {
		query += " AND id_producto = ?"
		args = append(args, idProducto)
	}
	if idDeposito != nil {
		query += " AND id_deposito = ?"
		args = append(args, idDeposito)
	}

	return strings.Replace(query, " AND ", " WHERE ", 1), args
}

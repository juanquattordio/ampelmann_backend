package receta

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.Receta {
	repo := Repository{
		db: db,
	}
	return &repo
}

func (r *Repository) CreateHeaderReceta(tx *sqlx.Tx, receta *entities.RecetaHeader) (int64, error) {
	var (
		idHeader int64
		stmt     *sql.Stmt
		err      error
	)
	if tx == nil {
		stmt, err = r.db.Prepare(insertRecetaHeader)
	} else {
		stmt, err = tx.Prepare(insertRecetaHeader)
	}
	result, err := stmt.Exec(&receta.PasoPaso, &receta.IdProductoFinal, &receta.LitrosFinales)
	if err != nil {
		return 0, errors.NewInternalServer(fmt.Sprintf("Fallo al crear receta de producto final id %d", *receta.IdProductoFinal))
	}
	idHeader, _ = result.LastInsertId()
	return idHeader, nil
}

func (r *Repository) CreateLineReceta(tx *sqlx.Tx, idHeader int64, idInsumo *int64, cantidad *float64, observaciones *string) error {
	var (
		stmt *sql.Stmt
		err  error
	)
	if tx == nil {
		stmt, err = r.db.Prepare(insertRecetaLine)
	} else {
		stmt, err = tx.Prepare(insertRecetaLine)
	}
	_, err = stmt.Exec(idHeader, &idInsumo, &cantidad, &observaciones)
	if err != nil {
		return errors.NewInternalServer(fmt.Sprintf("Fallo al asociar insumo id %d con receta header id %d", *idInsumo, idHeader))
	}
	return nil
}

func (r *Repository) CreateReceta(ctx context.Context, header *entities.RecetaHeader) error {
	tx := r.db.MustBegin()
	var err error
	var idHeader int64

	// inserta header en tabla
	if idHeader, err = r.CreateHeaderReceta(tx, header); err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			return errors.NewInternalServer("Fallo en el rollback de la transacción")
		}
		return errors.NewInternalServer("Fallo en crear header receta. Se hace rollback")
	}
	header.IdHeader = idHeader

	// por cada linea del comprobante se actualiza stocks e inserta linea de comprobante en tabla
	for _, linea := range header.Ingredientes {
		if err = r.CreateLineReceta(tx, idHeader, &linea.IdInsumo, &linea.Cantidad, &linea.Observaciones); err != nil {
			break
		}
	}
	if err != nil {
		if errRollBack := tx.Rollback(); errRollBack != nil {
			return errors.NewInternalServer("Fallo en el rollback de la transacción")
		}
		return errors.NewInternalServer("Fallo en crear lineas de Receta. Se hace rollback")
	}

	if err = tx.Commit(); err != nil {
		return errors.NewInternalServer("Fallo en el commit de la transacción")
	}
	return nil
}

//func (r *Repository) UpdateReceta(tx *sqlx.Tx, idInsumo *int64, idDeposito *int64, cantidad float64) error {
//	var err error
//	if tx == nil {
//		_, err = r.db.Query(updateReceta, &idDeposito, &idInsumo, &cantidad)
//	} else {
//		_, err = tx.Query(updateReceta, &idDeposito, &idInsumo, &cantidad)
//		if err != nil {
//			tx.Rollback()
//			return errors.NewInternalServer("Fallo al actualizar stock")
//		}
//	}
//
//	return err
//}

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

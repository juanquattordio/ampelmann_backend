package receta

import (
	"context"
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
		return 0, errors.NewInternalServer(fmt.Sprintf("Fallo al crear recetaHeader de producto final id %d", *receta.IdProductoFinal))
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
		return errors.NewInternalServer(fmt.Sprintf("Fallo al asociar insumo id %d con recetaHeader header id %d", *idInsumo, idHeader))
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
		return errors.NewInternalServer("Fallo en crear header recetaHeader. Se hace rollback")
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

func (r *Repository) DeleteReceta(tx *sqlx.Tx, idReceta int64) error {
	//stmt, err := r.db.Prepare(deleteRecetaHeader)
	_, err := tx.Query(deleteRecetaHeader, idReceta)
	if err != nil {
		tx.Rollback()
		return errors.NewInternalServer("Fallo al eliminar receta header")
	}
	//fmt.Sprintf("%v", rows.r)
	if _, err = tx.Query(deleteRecetaIngredientes, idReceta); err != nil {
		tx.Rollback()
		return errors.NewInternalServer("Fallo al eliminar ingredientes de receta receta header")
	}
	return nil
}

func (r *Repository) UpdateReceta(idReceta int64, receta *entities.RecetaHeader) error {
	tx := r.db.MustBegin()

	if err := r.DeleteReceta(tx, idReceta); err != nil {
		return errors.NewInternalServer("Fallo al actualizar receta header")
	}

	if _, err := tx.Query(reinsertRecetaHeader, receta.IdHeader, receta.PasoPaso, receta.IdProductoFinal, receta.LitrosFinales); err != nil {
		tx.Rollback()
		return errors.NewInternalServer("Fallo al actualizar receta header")
	}
	for _, ingrediente := range receta.Ingredientes {
		if _, err := tx.Query(insertRecetaLine, receta.IdHeader, ingrediente.IdInsumo, ingrediente.Cantidad, ingrediente.Observaciones); err != nil {
			tx.Rollback()
			return errors.NewInternalServer(fmt.Sprintf("Fallo al actualizar ingrediente insumo %d de receta %d", ingrediente.IdInsumo, receta.IdHeader))
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.NewInternalServer("Fallo en el commit de la transacción")
	}

	return nil
}

func (r *Repository) Search(idReceta *int64) (*entities.RecetaHeader, error) {
	rows, err := r.db.Queryx(getRecetaDetails, idReceta)
	if err != nil {
		return nil, err
	}

	var ingredientes []entities.Ingredientes

	for rows.Next() {
		i := entities.Ingredientes{}
		_ = rows.Scan(&i.IdInsumo, &i.UnidadMedida, &i.Cantidad, &i.Observaciones)
		ingredientes = append(ingredientes, i)
	}
	rh := recetaHeader{}
	rows.StructScan(&rh)

	recetaResult := rh.toEntity()
	recetaResult.Ingredientes = ingredientes

	return recetaResult, nil
}

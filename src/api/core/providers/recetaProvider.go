package providers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Receta interface {
	CreateReceta(ctx context.Context, movimiento *entities.RecetaHeader) error
	Search(idReceta *int64) (*entities.RecetaHeader, error)
	UpdateReceta(idReceta int64, receta *entities.RecetaHeader) error
	DeleteReceta(tx *sqlx.Tx, idReceta int64) error
	CreateHeaderReceta(tx *sqlx.Tx, receta *entities.RecetaHeader) (int64, error)
	CreateLineReceta(tx *sqlx.Tx, idHeader int64, idInsumo *int64, cantidad *float64, observaciones *string) error
}

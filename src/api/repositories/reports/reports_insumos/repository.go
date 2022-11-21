package reports_insumos

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) providers.ReportsProvider {
	repo := Repository{
		db: db,
	}
	return &repo
}

var LastIdInsumo int64

func (r *Repository) GetStockInsumosDesactivados() ([]entities.Insumo, error) {
	var dbInsumos []insumo

	err := r.db.Select(&dbInsumos, getStockInsumosDesactivados)

	if err != nil {
		return nil, err
	}

	insumoResult := toEntities(dbInsumos)
	return insumoResult, nil
}

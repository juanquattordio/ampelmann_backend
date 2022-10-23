package batch

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

const (
	resource = "batch"
)

const (
	saveScriptMySQL = "INSERT INTO Produccion_Batch(id_receta, fecha, litros_producidos) VALUES(?, ?, ?) "
	lastBatchId     = "SELECT MAX(id_batch) FROM Produccion_Batch"
	batchByBatch    = "SELECT id_batch, id_receta, fecha, multiplo_receta FROM Produccion_Batch"
	batchByReceta   = "SELECT id_batch, id_receta, fecha, multiplo_receta FROM Produccion_Batch WHERE id_receta = ?"
)

type batch struct {
	ID               int64     `db:"id_batch"`
	IdReceta         int64     `db:"id_receta"`
	Fecha            time.Time `db:"fecha"`
	LitrosProducidos float64   `db:"litros_producidos"`
}

func (dbItem batch) toEntity() *entities.Batch {
	return &entities.Batch{
		IdBatch:          dbItem.ID,
		IdReceta:         dbItem.IdReceta,
		Fecha:            dbItem.Fecha,
		LitrosProducidos: dbItem.LitrosProducidos,
	}
}

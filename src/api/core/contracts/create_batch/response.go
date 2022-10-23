package create_batch

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

type Response struct {
	IdBatch          int64     `form:"id_batch" json:"id_batch"`
	IdReceta         int64     `form:"id_receta" json:"id_receta"`
	Fecha            time.Time `form:"fecha" json:"fecha"`
	LitrosProducidos float64   `form:"litros_producidos" json:"litros_producidos"`
}

func NewResponse(batch *entities.Batch) *Response {
	return &Response{
		IdBatch:          batch.IdBatch,
		IdReceta:         batch.IdReceta,
		Fecha:            batch.Fecha,
		LitrosProducidos: batch.LitrosProducidos,
	}
}

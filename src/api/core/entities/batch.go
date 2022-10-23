package entities

import "time"

type Batch struct {
	IdBatch          int64
	IdReceta         int64
	Fecha            time.Time
	LitrosProducidos float64
}

func NewBatch(idBatch int64, idReceta int64, fecha time.Time) *Batch {
	op := &Batch{
		IdBatch:  idBatch,
		IdReceta: idReceta,
		Fecha:    fecha,
	}

	return op
}

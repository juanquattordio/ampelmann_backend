package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Batch interface {
	CreateBatch(batch *entities.Batch) error
	GetLastBacth() (int64, error)
	DeleteBatch(tx *sqlx.Tx, idBatch int64) error
}

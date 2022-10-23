package providers

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Batch interface {
	CreateBatch(batch *entities.Batch) error
	DeleteBatch(tx *sqlx.Tx, idBatch int64) error
}

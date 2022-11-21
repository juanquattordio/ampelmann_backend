package insumos_reports

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	GetStockInsumosDesactivados() ([]entities.Insumo, error)
}

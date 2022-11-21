package reports

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

type UseCase interface {
	GetStockInsumosDesactivados() ([]entities.Insumo, error)
	GetStockProductosDesactivados() ([]entities.ProductoFinal, error)
	GetClientesDesactivados() ([]entities.Cliente, error)
	GetFacturacionBetweenDates(dateTo, dateFrom time.Time) (float64, error)
}

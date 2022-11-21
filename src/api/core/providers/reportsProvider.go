package providers

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

type ReportsProvider interface {
	GetStockInsumosDesactivados() ([]entities.Insumo, error)
	GetStockProductosDesactivados() ([]entities.ProductoFinal, error)
	GetClientesDesactivados() ([]entities.Cliente, error)
}

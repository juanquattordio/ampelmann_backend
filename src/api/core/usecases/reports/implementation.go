package reports

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"time"
)

type Implementation struct {
	ReportsProvider providers.ReportsProvider
}

func (uc *Implementation) GetStockInsumosDesactivados() ([]entities.Insumo, error) {
	stock, err := uc.ReportsProvider.GetStockInsumosDesactivados()
	if err != nil {
		return nil, err
	}
	return stock, err
}

func (uc *Implementation) GetStockProductosDesactivados() ([]entities.ProductoFinal, error) {
	stock, err := uc.ReportsProvider.GetStockProductosDesactivados()
	if err != nil {
		return nil, err
	}
	return stock, err
}

func (uc *Implementation) GetClientesDesactivados() ([]entities.Cliente, error) {
	clientes, err := uc.ReportsProvider.GetClientesDesactivados()
	if err != nil {
		return nil, err
	}
	return clientes, err
}

func (uc *Implementation) GetFacturacionBetweenDates(dateTo, dateFrom time.Time) (float64, error) {
	totalAmount, err := uc.ReportsProvider.GetFacturacionBetweenDates(dateTo, dateFrom)
	if err != nil {
		return 0, err
	}
	return totalAmount, err
}

package insumos_reports

import (
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ReportsProvider providers.ReportsProvider
}

var (
	ErrNotFoundInsumo   = goErrors.New("insumo not found")
	ErrNotFoundDeposito = goErrors.New("deposito not found")
	ErrInternal         = goErrors.New("internal error")
)

func (uc *Implementation) GetStockInsumosDesactivados() ([]entities.Insumo, error) {
	stock, err := uc.ReportsProvider.GetStockInsumosDesactivados()
	if err != nil {
		return nil, err
	}
	return stock, err

}

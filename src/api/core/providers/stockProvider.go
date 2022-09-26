package providers

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

type Stock interface {
	GetStockInsumo(idInsumo *int64, idDeposito *int64) (float64, error)
	GetStockDeposito(idDeposito *int64) ([]entities.Insumo, error)
	//Update(insumo *entities.Insumo) error
}

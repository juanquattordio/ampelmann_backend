package get_stock

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	InsumoProvider   providers.Insumo
	DepositoProvider providers.Deposito
	//InsumoRepo    insumo.Repository
	//DepositoRepo  deposito.Repository
	StockProvider providers.Stock
}

var (
	ErrNotFoundInsumo   = goErrors.New("insumo not found")
	ErrNotFoundDeposito = goErrors.New("deposito not found")
	ErrInternal         = goErrors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, idInsumo *int64, idDeposito *int64) (*entities.Insumo, *entities.Deposito, float64, error) {

	// valida que exista la entidad a trabajar
	insumoDB, err := uc.InsumoProvider.Search(idInsumo, nil)
	if insumoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, nil, 0, ErrNotFoundInsumo
	}
	depositoDB := new(entities.Deposito)
	if idDeposito != nil {
		depositoDB, err = uc.DepositoProvider.Search(idDeposito, nil)
		if depositoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
			return nil, nil, 0, ErrNotFoundDeposito
		}
	}

	stock, err := uc.StockProvider.GetStock(idInsumo, idDeposito)
	return insumoDB, depositoDB, stock, err

	//if idDeposito != nil {
	//	var stockByInsumoList []stockByInsumo
	//	var err error
	//	stockByInsumoList, err = uc.StockProvider.GetStockDeposito(idDeposito)
	//	return stockByInsumoList, err
	//}

}

type stockByInsumo struct {
	idInsumo int
	nombre   string
	stock    float64
}

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
	StockProvider    providers.Stock
}

var (
	ErrNotFoundInsumo   = goErrors.New("insumo not found")
	ErrNotFoundDeposito = goErrors.New("deposito not found")
	ErrInternal         = goErrors.New("internal error")
)

func (uc *Implementation) GetStockByInsumo(ctx context.Context, idInsumo *int64, idDeposito *int64) (*entities.Insumo, *entities.Deposito, error) {

	// valida que exista la entidad a trabajar
	insumoDB, err := uc.InsumoProvider.Search(idInsumo, nil)
	if insumoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotFoundInsumo
	}
	depositoDB := new(entities.Deposito)
	if idDeposito != nil {
		depositoDB, err = uc.DepositoProvider.Search(idDeposito, nil)
		if depositoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrNotFoundDeposito
		}
	}

	stock, err := uc.StockProvider.GetStockInsumo(idInsumo, idDeposito)
	insumoDB.Stock = stock
	return insumoDB, depositoDB, err

}

func (uc *Implementation) GetStockByDeposito(ctx context.Context, idDeposito *int64) (*entities.Deposito, []entities.Insumo, error) {
	//TODO Nunca va a llegar un idDeposito nil, se chequea con el validate en el handler.
	depositoDB, err := uc.DepositoProvider.Search(idDeposito, nil)
	if depositoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotFoundDeposito
	}

	insumos, err := uc.StockProvider.GetStockDeposito(ctx, idDeposito)
	if err != nil {
		return nil, nil, err
	}
	return depositoDB, insumos, err
}

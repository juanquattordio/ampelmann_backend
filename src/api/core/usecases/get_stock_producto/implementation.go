package get_stock_producto

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ProductoProvider providers.ProductoFinal
	DepositoProvider providers.Deposito
	StockProvider    providers.StockProducto
}

var (
	ErrNotFoundProducto = goErrors.New("producto not found")
	ErrNotFoundDeposito = goErrors.New("deposito not found")
	ErrInternal         = goErrors.New("internal error")
)

func (uc *Implementation) GetStockByProducto(ctx context.Context, idProducto *int64, idDeposito *int64) (*entities.ProductoFinal, *entities.Deposito, error) {

	// valida que exista la entidad a trabajar
	productoDB, err := uc.ProductoProvider.Search(idProducto, nil)
	if productoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotFoundProducto
	}
	depositoDB := new(entities.Deposito)
	if idDeposito != nil {
		depositoDB, err = uc.DepositoProvider.Search(idDeposito, nil)
		if depositoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
			return nil, nil, ErrNotFoundDeposito
		}
	}

	stock, err := uc.StockProvider.GetStockProducto(idProducto, idDeposito)
	productoDB.Stock = stock
	return productoDB, depositoDB, err

}

func (uc *Implementation) GetStockByDeposito(ctx context.Context, idDeposito *int64) (*entities.Deposito, []entities.ProductoFinal, error) {
	//TODO Nunca va a llegar un idDeposito nil, se chequea con el validate en el handler.
	depositoDB, err := uc.DepositoProvider.Search(idDeposito, nil)
	if depositoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, nil, ErrNotFoundDeposito
	}

	productos, err := uc.StockProvider.GetStockDeposito(ctx, idDeposito)
	if err != nil {
		return nil, nil, err
	}
	return depositoDB, productos, err
}

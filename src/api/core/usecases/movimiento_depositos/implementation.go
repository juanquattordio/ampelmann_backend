package movimiento_depositos

import (
	"context"
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	StockProvider    providers.Stock
	DepositoProvider providers.Deposito
}

var (
	ErrNotFound          = goErrors.New("insumo not found")
	ErrInsufficientStock = goErrors.New("the stock is insufficient. Operation cancelled.")
	ErrInternal          = goErrors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, req movimiento_depositos.Request) (*entities.MovimientoHeader, error) {

	// verifico stock disponible en origen
	stockInsumos, err := uc.StockProvider.GetStockDeposito(ctx, req.IdDepositoOrigen)
	// Si el deposito no está cargado, falla
	if err != nil || goErrors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	stockSuficiente, msgStockInsuficiente := validarStockOrigen(req.Insumos, stockInsumos)
	if !stockSuficiente {
		return nil, errors.NewInternalServer(msgStockInsuficiente)
	}

	// verifico deposito de destino
	if deposito, err := uc.DepositoProvider.Search(req.IdDepositoDestino, nil); err != nil ||
		deposito.Status == constants.Desactivo {
		return nil, goErrors.New(fmt.Sprintf("Deposito destino no existe o está desactivado"))
	}

	// Crea un movimiento que ejecuta Updates de stocks en cada depósito
	var causaMovimiento string
	if req.CausaMovimiento == nil || *req.CausaMovimiento == "" {
		causaMovimiento = "Ajuste stock"
	} else {
		causaMovimiento = *req.CausaMovimiento
	}
	movimiento := entities.NewMovimientoDeposito(*req.IdDepositoOrigen, *req.IdDepositoDestino, toEntities(req.Insumos), causaMovimiento)
	err = uc.StockProvider.MovimientoDepositos(ctx, movimiento)

	if err != nil {
		return nil, err
	}

	return movimiento, nil
}
func validarStockOrigen(movimientos []movimiento_depositos.Insumo, stockDeposito []entities.Insumo) (stockSuficiente bool, msg string) {
	stockSuficiente = false
	var msgStockInsuficiente string
	// Todo Se podría hacer una búsqueda ordenada o algo más eficiente
	for _, lineaMov := range movimientos {
		stockSuficiente = false
		msgStockInsuficiente += fmt.Sprintf("Stock insuficiente id_producto: %d en en deposito de origen de este movimiento.",
			*lineaMov.IdInsumo)
		for i := 0; i < len(stockDeposito); i++ {
			if *lineaMov.IdInsumo == stockDeposito[i].IdInsumo {
				msgStockInsuficiente += fmt.Sprintf("Stock insuficiente id_producto: %d con stock %.2f para este movimiento.",
					stockDeposito[i].IdInsumo, stockDeposito[i].Stock)
				if stockDeposito[i].Stock >= *lineaMov.Cantidad {
					stockSuficiente = true
					stockDeposito[i].Stock = stockDeposito[i].Stock - *lineaMov.Cantidad
					msgStockInsuficiente = ""
				}
				break
			}
		}
		if !stockSuficiente {
			break
		}
	}
	return stockSuficiente, msgStockInsuficiente
}

func toEntities(insumosRequest []movimiento_depositos.Insumo) []entities.MovimientoLine {
	var lineas []entities.MovimientoLine
	for _, insumo := range insumosRequest {
		line := new(entities.MovimientoLine)
		line.IdInsumo = *insumo.IdInsumo
		line.Cantidad = *insumo.Cantidad
		line.Obseraciones = insumo.Obseraciones
		lineas = append(lineas, *line)
	}
	return lineas
}

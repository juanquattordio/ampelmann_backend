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
	StockProvider         providers.Stock
	StockProductoProvider providers.StockProducto
	DepositoProvider      providers.Deposito
}

var (
	ErrNotFound          = goErrors.New("insumo not found")
	ErrInsufficientStock = goErrors.New("the stock is insufficient. Operation cancelled.")
	ErrInternal          = goErrors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, req movimiento_depositos.Request) (*entities.MovimientoHeader, error) {
	var (
		movimientoLines []entities.MovimientoLine
		tipoArticulo    string
	)
	if req.Productos == nil {
		// verifico stock disponible en origen
		stockInsumos, err := uc.StockProvider.GetStockDeposito(ctx, &req.IdDepositoOrigen)
		// Si el deposito no está cargado, falla
		if err != nil || goErrors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		if stockInsumos == nil {
			return nil, ErrInsufficientStock
		}
		stockSuficiente, msgStockInsuficiente := validarStockInsumosOrigen(req.Insumos, stockInsumos)
		if !stockSuficiente {
			return nil, errors.NewInternalServer(msgStockInsuficiente)
		}
		movimientoLines = toEntities(req.Insumos)
		tipoArticulo = constants.Insumos
	} else {
		// verifico stock disponible en origen
		stockProductos, err := uc.StockProductoProvider.GetStockDeposito(ctx, &req.IdDepositoOrigen)
		// Si el deposito no está cargado, falla
		if err != nil || goErrors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		if stockProductos == nil {
			return nil, ErrInsufficientStock
		}
		stockSuficiente, msgStockInsuficiente := validarStockProductosOrigen(req.Productos, stockProductos)
		if !stockSuficiente {
			return nil, errors.NewInternalServer(msgStockInsuficiente)
		}
		movimientoLines = toEntities(req.Productos)
		tipoArticulo = constants.Productos

	}

	// verifico deposito de destino
	if deposito, err := uc.DepositoProvider.Search(&req.IdDepositoDestino, nil); err != nil ||
		deposito.Status == constants.Desactivo {
		return nil, goErrors.New(fmt.Sprintf("Deposito destino no existe o está desactivado"))
	}

	// Crea un movimiento que ejecuta Updates de stocks en cada depósito
	var causaMovimiento string
	if &req.CausaMovimiento == nil || req.CausaMovimiento == "" {
		causaMovimiento = "Ajuste stock"
	} else {
		causaMovimiento = req.CausaMovimiento
	}

	movimiento := entities.NewMovimientoDeposito(req.IdDepositoOrigen, req.IdDepositoDestino, movimientoLines, causaMovimiento)
	err := uc.StockProvider.MovimientoDepositos(ctx, movimiento, tipoArticulo)

	if err != nil {
		return nil, err
	}

	return movimiento, nil
}
func validarStockInsumosOrigen(movimientos []movimiento_depositos.Articulos, stockDeposito []entities.Insumo) (stockSuficiente bool, msg string) {
	stockSuficiente = false
	var msgStockInsuficiente string
	// Todo Se podría hacer una búsqueda ordenada o algo más eficiente
	for _, lineaMov := range movimientos {
		stockSuficiente = false
		for i := 0; i < len(stockDeposito); i++ {
			if *lineaMov.IdArticulo == stockDeposito[i].IdInsumo {
				msgStockInsuficiente = fmt.Sprintf("Stock insuficiente id_producto: %d con stock %.2f para este movimiento se necesitan %.2f.",
					stockDeposito[i].IdInsumo, stockDeposito[i].Stock, *lineaMov.Cantidad)
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

func validarStockProductosOrigen(movimientos []movimiento_depositos.Articulos, stockDeposito []entities.ProductoFinal) (stockSuficiente bool, msg string) {
	stockSuficiente = false
	var msgStockInsuficiente string
	// Todo Se podría hacer una búsqueda ordenada o algo más eficiente
	for _, lineaMov := range movimientos {
		stockSuficiente = false
		for i := 0; i < len(stockDeposito); i++ {
			if *lineaMov.IdArticulo == stockDeposito[i].Id {
				msgStockInsuficiente = fmt.Sprintf("Stock insuficiente id_producto: %d con stock %.2f para este movimiento se necesitan %.2f.",
					stockDeposito[i].Id, stockDeposito[i].Stock, *lineaMov.Cantidad)
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

func toEntities(insumosRequest []movimiento_depositos.Articulos) []entities.MovimientoLine {
	var lineas []entities.MovimientoLine
	for _, ingrediente := range insumosRequest {
		line := new(entities.MovimientoLine)
		line.IdArticulo = *ingrediente.IdArticulo
		line.Cantidad = *ingrediente.Cantidad
		line.Observaciones = ingrediente.Observaciones
		lineas = append(lineas, *line)
	}
	return lineas
}

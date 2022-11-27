package create_factura_venta

import (
	"context"
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura"
	movimientoDeposito "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/movimiento_depositos"
)

type Implementation struct {
	ClienteProvider           providers.Cliente
	DocumentoProvider         providers.Documento
	ProductoProvider          providers.ProductoFinal
	MovimientoDepositoUseCase movimiento_depositos.UseCase
}

func (uc *Implementation) Execute(ctx context.Context, request create_factura.RequestFacturaVenta) (*entities.FacturaVentaHeader, error) {
	// valida que exista el cliente
	clienteExists, err := uc.ClienteProvider.Search(request.IdCliente, nil)
	if clienteExists == nil || goErrors.Is(err, sql.ErrNoRows) {
		return nil, goErrors.New(fmt.Sprintf("id_cliente %d not found", *request.IdCliente))
	}
	// valida que existan los productos a facturar
	for _, linea := range request.Lineas {
		productoExists, err := uc.ProductoProvider.Search(linea.IdArticulo, nil)
		if err != nil || productoExists == nil || goErrors.Is(err, sql.ErrNoRows) {
			return nil, goErrors.New(fmt.Sprintf("id_producto %d not found", *linea.IdArticulo))
		}
	}

	lastIdFactura, err := uc.DocumentoProvider.LastFacturaVenta()

	// Verifica stock disponible en origen y crea un movimiento que ejecuta Updates de stocks en cada depósito
	idDepositoProductos := int64(3)
	if request.IdDepositoOrigen != 0 {
		idDepositoProductos = request.IdDepositoOrigen
	}
	causaMovimiento := fmt.Sprintf("FVC-%d", lastIdFactura+1)
	reqProductoMov := movimientoDeposito.Request{
		IdDepositoOrigen: idDepositoProductos,
		Productos:        parseToArticles(request.Lineas),
		CausaMovimiento:  causaMovimiento}
	_, err = uc.MovimientoDepositoUseCase.Execute(ctx, reqProductoMov)
	if err != nil {
		return nil, err
	}

	newFactura := entities.NewFacturaVenta(*request.IdCliente, toEntities(request.Lineas), request.Observaciones)
	if err = uc.DocumentoProvider.CreateFacturaVenta(newFactura); err != nil {
		return nil, goErrors.New(fmt.Sprintf("fallo en la creación de la factura"))
	}

	return newFactura, nil
}

func toEntities(linesRequest []create_factura.FacturaLine) []entities.FacturaLine {
	var lineas []entities.FacturaLine
	for _, lineReq := range linesRequest {
		line := new(entities.FacturaLine)
		line.IdArticulo = *lineReq.IdArticulo
		line.Cantidad = *lineReq.Cantidad
		line.PrecioUnitario = *lineReq.PrecioUnitario
		line.Observaciones = lineReq.Observaciones
		lineas = append(lineas, *line)
	}
	return lineas
}

func parseToArticles(lineas []create_factura.FacturaLine) []movimientoDeposito.Articulos {
	articulos := make([]movimientoDeposito.Articulos, len(lineas))
	for i := range lineas {
		articulos[i].IdLinea = int64(i + 1)
		articulos[i].IdArticulo = lineas[i].IdArticulo
		articulos[i].Cantidad = lineas[i].Cantidad
	}
	return articulos
}

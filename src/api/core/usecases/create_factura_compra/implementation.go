package create_factura_compra

import (
	"context"
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"time"
)

type Implementation struct {
	ProveedorProvider providers.Proveedor
	DocumentoProvider providers.Documento
	InsumoProvider    providers.Insumo
	StockProvider     providers.Stock
}

func (uc *Implementation) Execute(ctx context.Context, request create_factura.RequestFacturaCompra) (*entities.FacturaCompraHeader, error) {
	// valida que exista el proveedor
	proveedorExists, err := uc.ProveedorProvider.Search(request.IdProveedor, nil)
	if proveedorExists == nil || goErrors.Is(err, sql.ErrNoRows) {
		return nil, goErrors.New(fmt.Sprintf("id_proveedor %d not found", *request.IdProveedor))
	}
	// valida que existan los insumos comprados
	for _, linea := range request.Lineas {
		productoExists, err := uc.InsumoProvider.Search(linea.IdArticulo, nil)
		if err != nil || productoExists == nil || goErrors.Is(err, sql.ErrNoRows) {
			return nil, goErrors.New(fmt.Sprintf("id_insumo %d not found", *linea.IdArticulo))
		}
		// Carga historial de precio por insumo y proveedor
		err = uc.ProveedorProvider.UpdateHistorialPrecioInsumo(request.IdProveedor, linea.IdArticulo,
			linea.PrecioUnitario, time.Time(request.FechaOrigen), "")
		if err != nil {
			return nil, goErrors.New(fmt.Sprintf("Error al actualizar historial de precios"))
		}
	}

	lastIdFactura, err := uc.DocumentoProvider.LastFacturaCompra()

	// Crea un movimiento que ejecuta Updates de stocks en cada depósito
	idDepositoInsumos := int64(2)
	if request.IdDepositoDestino != 0 {
		idDepositoInsumos = request.IdDepositoDestino
	}
	causaMovimiento := fmt.Sprintf("FCP-%d", lastIdFactura+1)
	movimiento := entities.NewMovimientoDeposito(0, idDepositoInsumos, parseToMovLines(request.Lineas), causaMovimiento)
	if err = uc.StockProvider.MovimientoDepositos(ctx, movimiento, constants.Insumos); err != nil {
		return nil, goErrors.New(fmt.Sprintf("fallo en la creación del movimiento de insumos por compra"))
	}

	newFactura := entities.NewFacturaCompra(*request.IdProveedor, *request.IdFacturaProveedor, time.Time(request.FechaOrigen), toEntities(request.Lineas))

	if err = uc.DocumentoProvider.CreateFacturaCompra(newFactura); err != nil {
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

func parseToMovLines(linesRequest []create_factura.FacturaLine) []entities.MovimientoLine {
	var lineas []entities.MovimientoLine
	for _, lineReq := range linesRequest {
		line := new(entities.MovimientoLine)
		line.IdArticulo = *lineReq.IdArticulo
		line.Cantidad = *lineReq.Cantidad
		line.Observaciones = lineReq.Observaciones
		lineas = append(lineas, *line)
	}
	return lineas
}

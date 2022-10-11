package create_factura_compra

import (
	"context"
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura_compra"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"time"
)

type Implementation struct {
	ProveedorProvider providers.Proveedor
	DocumentoProvider providers.Documento
	InsumoProvider    providers.Insumo
	StockProvider     providers.Stock
}

var (
	ErrNotFoundProveedor = goErrors.New("proveedor not found")
	ErrNotFoundInsumo    = goErrors.New("insumo not found")
	ErrInternal          = goErrors.New("internal error")
	ErrFieldsEmpty       = goErrors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_factura_compra.Request) (*entities.FacturaCompraHeader, error) {
	// valida que exista el proveedor
	proveedorExists, err := uc.ProveedorProvider.Search(request.IdProveedor, nil)
	if proveedorExists == nil || goErrors.Is(err, sql.ErrNoRows) {
		return nil, goErrors.New(fmt.Sprintf("id_proveedor %d not found", *request.IdProveedor))
	}
	// valida que existan los insumos comprados
	for _, linea := range request.Lineas {
		productoExists, err := uc.InsumoProvider.Search(linea.IdInsumo, nil)
		if err != nil || productoExists == nil || goErrors.Is(err, sql.ErrNoRows) {
			return nil, goErrors.New(fmt.Sprintf("id_proveedor %d not found", linea.IdInsumo))
		}
		// Carga historial de precio por insumo y proveedor
		err = uc.ProveedorProvider.UpdateHistorialPrecioInsumo(request.IdProveedor, linea.IdInsumo,
			linea.PrecioUnitario, time.Time(request.FechaOrigen), "")
		if err != nil {
			return nil, goErrors.New(fmt.Sprintf("Error al actualizar historial de precios"))
		}
	}

	newFactura := entities.NewFacturaCompra(*request.IdProveedor, *request.IdFacturaProveedor, time.Time(request.FechaOrigen), toEntities(request.Lineas))

	if err = uc.DocumentoProvider.CreateFacturaCompra(newFactura); err != nil {
		return nil, goErrors.New(fmt.Sprintf("fallo en la creación de la factura"))
	}

	// Crea un movimiento que ejecuta Updates de stocks en cada depósito
	idDepositoInsumos := int64(2)
	movimiento := entities.NewMovimientoDeposito(0, idDepositoInsumos, parseToMovLines(request.Lineas))
	if err = uc.StockProvider.MovimientoDepositos(ctx, movimiento); err != nil {
		return nil, goErrors.New(fmt.Sprintf("fallo en la creación del movimiento de insumos por compra"))
	}

	return newFactura, nil
}

func toEntities(linesRequest []create_factura_compra.FacturaCompraLine) []entities.FacturaCompraLine {
	var lineas []entities.FacturaCompraLine
	for _, lineReq := range linesRequest {
		line := new(entities.FacturaCompraLine)
		line.IdInsumo = *lineReq.IdInsumo
		line.Cantidad = *lineReq.Cantidad
		line.PrecioUnitario = *lineReq.PrecioUnitario
		line.Obseraciones = lineReq.Obseraciones
		lineas = append(lineas, *line)
	}
	return lineas
}

func parseToMovLines(linesRequest []create_factura_compra.FacturaCompraLine) []entities.MovimientoLine {
	var lineas []entities.MovimientoLine
	for _, lineReq := range linesRequest {
		line := new(entities.MovimientoLine)
		line.IdInsumo = *lineReq.IdInsumo
		line.Cantidad = *lineReq.Cantidad
		line.Obseraciones = lineReq.Obseraciones
		lineas = append(lineas, *line)
	}
	return lineas
}

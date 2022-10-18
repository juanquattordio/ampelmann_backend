package entities

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"time"
)

type FacturaCompraHeader struct {
	IdFactura          int64
	Fecha              time.Time
	IdProveedor        int64
	IdFacturaProveedor string
	FechaOrigen        time.Time
	Lineas             []FacturaLine
	ImporteTotal       float64
	Status             string
}

type FacturaLine struct {
	IdLinea        int64
	IdArticulo     int64
	Cantidad       float64
	PrecioUnitario float64
	Observaciones  string
}

func NewFacturaCompra(idProveedor int64, idFacturaProveedor string, fechaOrigen time.Time, lineas []FacturaLine) *FacturaCompraHeader {
	facturaCompraHeader := &FacturaCompraHeader{
		Fecha:              time.Now().UTC(),
		IdProveedor:        idProveedor,
		IdFacturaProveedor: idFacturaProveedor,
		FechaOrigen:        fechaOrigen,
		Lineas:             lineas,
		ImporteTotal:       calcularImporteTotal(lineas),
		Status:             constants.Activo,
	}

	return facturaCompraHeader
}

func calcularImporteTotal(lineas []FacturaLine) float64 {
	var total float64
	for _, linea := range lineas {
		total += linea.Cantidad * linea.PrecioUnitario
	}
	return total
}

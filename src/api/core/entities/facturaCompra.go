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
	Lineas             []FacturaCompraLine
	ImporteTotal       float64
	Status             string
}

type FacturaCompraLine struct {
	IdLinea        int64
	IdInsumo       int64
	Cantidad       float64
	PrecioUnitario float64
	Obseraciones   string
}

func NewFacturaCompra(idProveedor int64, idFacturaProveedor string, fechaOrigen time.Time, lineas []FacturaCompraLine) *FacturaCompraHeader {
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

func calcularImporteTotal(lineas []FacturaCompraLine) float64 {
	var total float64
	for _, linea := range lineas {
		total += linea.Cantidad * linea.PrecioUnitario
	}
	return total
}

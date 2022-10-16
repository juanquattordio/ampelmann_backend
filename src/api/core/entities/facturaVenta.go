package entities

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"time"
)

type FacturaVentaHeader struct {
	IdFactura     int64
	Fecha         time.Time
	IdCliente     int64
	Lineas        []FacturaLine
	ImporteTotal  float64
	Observaciones string
	Status        string
}

func NewFacturaVenta(idCliente int64, lineas []FacturaLine, observaciones string) *FacturaVentaHeader {
	facturaVentaHeader := &FacturaVentaHeader{
		IdCliente:     idCliente,
		Fecha:         time.Now().UTC(),
		Lineas:        lineas,
		ImporteTotal:  calcularImporteTotal(lineas),
		Observaciones: observaciones,
		Status:        constants.Activo,
	}

	return facturaVentaHeader
}

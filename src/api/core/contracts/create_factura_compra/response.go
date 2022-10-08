package create_factura_compra

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

type Response struct {
	IdFactura          *int64              `form:"id_factura" json:"id_factura"`
	Fecha              time.Time           `form:"fecha" json:"fecha"`
	IdProveedor        *int64              `form:"id_proveedor" json:"id_proveedor"`
	IdFacturaProveedor *string             `form:"id_factura_proveedor" json:"id_factura_proveedor"`
	FechaOrigen        *time.Time          `form:"fecha_origen" json:"fecha_origen"`
	Lineas             []FacturaCompraLine `form:"lineas" json:"lineas"`
	ImporteTotal       *float64            `form:"importe_total" json:"importe_total"`
	Status             *string             `form:"status" json:"status"`
}

func NewResponse(factura *entities.FacturaCompraHeader) *Response {
	lines := make([]FacturaCompraLine, len(factura.Lineas))
	for i := range factura.Lineas {
		lines[i].IdLinea = &factura.Lineas[i].IdLinea
		lines[i].IdInsumo = &factura.Lineas[i].IdInsumo
		lines[i].Cantidad = &factura.Lineas[i].Cantidad
		lines[i].PrecioUnitario = &factura.Lineas[i].PrecioUnitario
		lines[i].Obseraciones = factura.Lineas[i].Obseraciones
	}
	return &Response{
		IdFactura:          &factura.IdFactura,
		Fecha:              factura.Fecha,
		IdProveedor:        &factura.IdProveedor,
		IdFacturaProveedor: &factura.IdFacturaProveedor,
		FechaOrigen:        &factura.FechaOrigen,
		Lineas:             lines,
		ImporteTotal:       &factura.ImporteTotal,
		Status:             &factura.Status,
	}
}

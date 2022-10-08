package create_factura_compra

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomFechaOrigen time.Time

type Request struct {
	IdProveedor        *int64              `form:"id_proveedor" json:"id_proveedor" binding:"required"`
	IdFacturaProveedor *string             `form:"id_factura_proveedor" json:"id_factura_proveedor"`
	FechaOrigen        CustomFechaOrigen   `form:"fecha_origen" json:"fecha_origen" binding:"required"`
	Lineas             []FacturaCompraLine `form:"lineas" json:"lineas" binding:"required"`
	Status             *string             `form:"status" json:"status" binding:""`
}

type FacturaCompraLine struct {
	IdLinea        *int64   `form:"id_linea" json:"id_linea"`
	IdInsumo       *int64   `form:"id_insumo" json:"id_insumo" binding:"required"`
	Cantidad       *float64 `form:"cantidad" json:"cantidad" binding:"required"`
	PrecioUnitario *float64 `form:"precio_unitario" json:"precio_unitario" binding:"required"`
	Obseraciones   string   `form:"observaciones" json:"observaciones"`
}

// Implement Marshaler and Unmarshaler interface
func (j *CustomFechaOrigen) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = CustomFechaOrigen(t)
	return nil
}

func (j CustomFechaOrigen) MarshalJSON() ([]byte, error) {
	miJson, err := json.Marshal(time.Time(j))
	return miJson, err
}

// Maybe a Format function for printing your date
func (j CustomFechaOrigen) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

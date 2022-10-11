package update_historial_precios_proveedor

import (
	"context"
	"time"
)

type UseCase interface {
	Execute(context context.Context, idProveedor *int64, idInsumo *int64, precioUnitario *float64, fecha time.Time, status string) error
}

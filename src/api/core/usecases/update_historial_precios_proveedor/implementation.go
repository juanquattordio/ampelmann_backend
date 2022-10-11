package update_historial_precios_proveedor

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"time"
)

type Implementation struct {
	ProveedorProvider providers.Proveedor
}

var (
	ErrNotFound = errors.New("proveedor not found")
	ErrInternal = errors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, idProveedor *int64, idInsumo *int64, precioUnitario *float64, fecha time.Time, status string) error {
	err := uc.ProveedorProvider.UpdateHistorialPrecioInsumo(idProveedor, idInsumo, precioUnitario, fecha, status)
	if err != nil {
		return err
	}
	return nil
}

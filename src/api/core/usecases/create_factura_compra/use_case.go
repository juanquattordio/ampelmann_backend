package create_factura_compra

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_factura.RequestFacturaCompra) (*entities.FacturaCompraHeader, error)
}

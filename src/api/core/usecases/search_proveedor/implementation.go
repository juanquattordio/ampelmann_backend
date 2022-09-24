package search_proveedor

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ProveedorProvider providers.Proveedor
}

var (
	ErrNotFound = errors.New("proveedor not found")
	ErrInternal = errors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, id *int64, cuit *string) (*entities.Proveedor, error) {
	proveedor, err := uc.ProveedorProvider.Search(id, cuit)
	if err != nil {
		return nil, err
	}
	return proveedor, nil
}

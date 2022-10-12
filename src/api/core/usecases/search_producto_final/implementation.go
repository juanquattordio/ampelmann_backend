package search_producto_final

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ProductoFinalProvider providers.ProductoFinal
}

var (
	ErrNotFound = errors.New("product not found")
	ErrInternal = errors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, id *int64, descripcion *string) (*entities.ProductoFinal, error) {
	product, err := uc.ProductoFinalProvider.Search(id, descripcion)
	if err != nil {
		return nil, err
	}
	return product, nil
}

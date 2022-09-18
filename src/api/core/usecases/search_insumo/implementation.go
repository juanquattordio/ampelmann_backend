package search_insumo

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	InsumoProvider providers.Insumo
}

var (
	ErrNotFound    = errors.New("insumo not found")
	ErrInternal    = errors.New("internal error")
	ErrWhCodeEmpty = errors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, id *int64, nombre *string) (*entities.Insumo, error) {
	insumo, err := uc.InsumoProvider.Search(id, nombre)
	if err != nil {
		return nil, err
	}
	return insumo, nil
}

package search_cliente

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ClienteProvider providers.Cliente
}

var (
	ErrNotFound = errors.New("cliente not found")
	ErrInternal = errors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, id *int64, cuit *string) (*entities.Cliente, error) {
	cliente, err := uc.ClienteProvider.Search(id, cuit)
	if err != nil {
		return nil, err
	}
	return cliente, nil
}

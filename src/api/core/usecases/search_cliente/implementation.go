package search_cliente

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ClienteProvider providers.Cliente
}

var (
	ErrNotFound    = errors.New("cliente not found")
	ErrInternal    = errors.New("internal error")
	ErrWhCodeEmpty = errors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request search_cliente.Request) (*entities.Cliente, error) {
	cliente, err := uc.ClienteProvider.Search(request)
	if err != nil {
		return &entities.Cliente{}, err
	}
	return &cliente, nil
}

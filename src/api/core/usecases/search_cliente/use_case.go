package search_cliente

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request search_cliente.Request) (*entities.Cliente, error)
}

package create_cliente

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_cliente.Request) (*entities.Cliente, error)
}

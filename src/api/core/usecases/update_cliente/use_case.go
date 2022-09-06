package create_cliente

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

func NewUseCase(clienteProvider providers.Cliente) *Implementation {
	return &Implementation{
		ClienteProvider: clienteProvider,
	}
}

type UseCase interface {
	Execute(context context.Context, request create_cliente.Request) (*entities.Cliente, error)
}

//type useCase struct {
//	repository cliente.Repository
//}

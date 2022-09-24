package create_deposito

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_deposito.Request) (*entities.Deposito, error)
}

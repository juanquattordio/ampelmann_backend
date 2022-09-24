package update_deposito

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id int64, request update_deposito.Request) (*entities.Deposito, error)
}

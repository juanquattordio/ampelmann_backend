package update_cliente

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id int64, request update_cliente.Request) (*entities.Cliente, error)
}

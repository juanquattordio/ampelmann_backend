package search_cliente

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id *int64, cuit *string) (*entities.Cliente, error)
}

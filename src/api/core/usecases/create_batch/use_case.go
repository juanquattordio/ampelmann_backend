package create_batch

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_batch"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_batch.Request) (*entities.Batch, error)
}

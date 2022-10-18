package create_receta

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, request create_receta.Request) (*entities.RecetaHeader, error)
}

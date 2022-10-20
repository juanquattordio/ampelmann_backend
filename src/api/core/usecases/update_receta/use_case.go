package update_receta

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id int64, request update_receta.Request) (*entities.RecetaHeader, error)
}

package update_insumo

import (
	"context"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type UseCase interface {
	Execute(context context.Context, id int64, request update_insumo.RequestUpdate) (*entities.Insumo, error)
}

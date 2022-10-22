package delete_receta

import (
	"context"
)

type UseCase interface {
	Execute(context context.Context, id int64) error
}

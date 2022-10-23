package delete_receta

import (
	"context"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	RecetaProvider providers.Receta
}

var ErrNotFound = goErrors.New("receta not found")

func (uc *Implementation) Execute(ctx context.Context, id int64) error {

	_, err := uc.RecetaProvider.Search(&id)
	// Si la receta no existe, falla
	if err != nil {
		return err
	}

	if err = uc.RecetaProvider.DeleteReceta(nil, id); err != nil {
		return err
	}

	return nil
}

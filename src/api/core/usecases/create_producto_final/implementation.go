package create_producto_final

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ProductoFinalProvider providers.ProductoFinal
}

var (
	ErrNotFound  = goErrors.New("producto final not found")
	ErrDuplicate = goErrors.New("name already exists. Operation cancelled.")
	ErrInternal  = goErrors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, request create_producto_final.Request) (*entities.ProductoFinal, error) {
	productoFinalExists, err := uc.ProductoFinalProvider.Search(nil, request.Descripcion)
	if productoFinalExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
	}

	newProductoFinal := entities.NewProductoFinal(*request.Descripcion)

	lastId, err := uc.ProductoFinalProvider.Save(*newProductoFinal)
	if err != nil {
		return nil, err
	}
	newProductoFinal.Id = lastId
	return newProductoFinal, nil
}

package create_producto_final

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Implementation struct {
	ProductoFinalProvider providers.ProductoFinal
}

var (
	ErrNotFound         = goErrors.New("producto final not found")
	ErrDuplicate        = goErrors.New("name already exists. Operation cancelled.")
	ErrDisavailableUnit = goErrors.New("unit of measurement disavailable. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_producto_final.Request) (*entities.ProductoFinal, error) {
	productoFinalExists, err := uc.ProductoFinalProvider.Search(nil, request.Descripcion)
	if productoFinalExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
	}

	if !isValidUnidad(*request.Unidad) {
		return nil, ErrDisavailableUnit
	}

	newProductoFinal := entities.NewProductoFinal(*request.Descripcion, *request.Unidad, 0)

	lastId, err := uc.ProductoFinalProvider.Save(*newProductoFinal)
	if err != nil {
		return nil, err
	}
	newProductoFinal.Id = lastId
	return newProductoFinal, nil
}

func isValidUnidad(unidad string) bool {
	unidad = strings.ToLower(unidad)
	return unidad == constants.LTS ||
		unidad == constants.UN
}

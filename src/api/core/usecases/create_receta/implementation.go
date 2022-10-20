package create_receta

import (
	"context"
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ProductoProvider providers.ProductoFinal
	InsumoProvider   providers.Insumo
	RecetaProvider   providers.Receta
}

var (
	ErrInsumoNotFound   = goErrors.New("ingrediente not found")
	ErrProductoNotFound = goErrors.New("producto final not found")
)

func (uc *Implementation) Execute(ctx context.Context, req create_receta.Request) (*entities.RecetaHeader, error) {

	// verifico exista el producto final
	_, err := uc.ProductoProvider.Search(req.IdProductoFinal, nil)
	// Si el producto final no está cargado, falla
	if err != nil || goErrors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("producto final id %d not found", *req.IdProductoFinal)
	}
	// verifico los insumos que utiliza y toma la unidad de medida de cada uno
	insumo := new(entities.Insumo)
	for i, ingrediente := range req.Ingredientes {
		if insumo, err = uc.InsumoProvider.Search(ingrediente.IdInsumo, nil); err != nil || goErrors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("insumo id %d not found", *ingrediente.IdInsumo)
		}
		req.Ingredientes[i].UnidadMedida = insumo.Unidad
	}

	receta := entities.NewReceta(0, *req.DetallePasoPaso, req.IdProductoFinal, toEntities(req.Ingredientes), *req.LitrosFinales)
	err = uc.RecetaProvider.CreateReceta(ctx, receta)

	if err != nil {
		return nil, err
	}

	return receta, nil
}

func toEntities(ingredientesRequest []create_receta.Ingredientes) []entities.Ingredientes {
	var lineas []entities.Ingredientes
	for _, ingrediente := range ingredientesRequest {
		line := new(entities.Ingredientes)
		line.IdInsumo = *ingrediente.IdInsumo
		line.UnidadMedida = ingrediente.UnidadMedida
		line.Cantidad = *ingrediente.Cantidad
		line.Observaciones = ingrediente.Observaciones
		lineas = append(lineas, *line)
	}
	return lineas
}

package update_receta

import (
	"context"
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	RecetaProvider   providers.Receta
	ProductoProvider providers.ProductoFinal
	InsumoProvider   providers.Insumo
}

func (uc *Implementation) Execute(ctx context.Context, id int64, req update_receta.Request) (*entities.RecetaHeader, error) {
	if err := uc.validateRequestUpdate(id, &req); err != nil {
		return nil, err
	}

	recetaToUpdate := entities.NewReceta(id, *req.DetallePasoPaso, req.IdProductoFinal, toEntities(req.Ingredientes), *req.LitrosFinales)

	if err := uc.prepareToUpdate(recetaToUpdate); err != nil {
		return nil, err
	}

	if err := uc.RecetaProvider.UpdateReceta(id, recetaToUpdate); err != nil {
		return nil, err
	}

	return recetaToUpdate, nil
}

func (uc *Implementation) validateRequestUpdate(idReceta int64, request *update_receta.Request) error {
	// verifico exista el producto final
	_, err := uc.ProductoProvider.Search(request.IdProductoFinal, nil)
	// Si el producto final no está cargado, falla
	if err != nil || goErrors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("producto final id %d not found", *request.IdProductoFinal)
	}
	// verifico los insumos que utiliza y toma la unidad de medida de cada uno
	insumo := new(entities.Insumo)
	for i, ingrediente := range request.Ingredientes {
		if insumo, err = uc.InsumoProvider.Search(ingrediente.IdInsumo, nil); err != nil || goErrors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("insumo id %d not found", *ingrediente.IdInsumo)
		}
		request.Ingredientes[i].UnidadMedida = insumo.Unidad
	}
	// verifico la receta a actualizar
	_, err = uc.RecetaProvider.Search(&idReceta)
	if err != nil {
		return err
	}
	return nil
}

func (uc *Implementation) prepareToUpdate(recetaToUpdate *entities.RecetaHeader) error {
	// compara contra los valores de la receta de la BD y actualiza según necesita.
	recetaBD, err := uc.RecetaProvider.Search(&recetaToUpdate.IdHeader)
	if err != nil {
		return err
	}

	if recetaToUpdate.PasoPaso == "" {
		recetaToUpdate.PasoPaso = recetaBD.PasoPaso
	}
	if recetaToUpdate.IdProductoFinal == nil {
		recetaToUpdate.IdProductoFinal = recetaBD.IdProductoFinal
	}
	if recetaToUpdate.Ingredientes == nil {
		recetaToUpdate.Ingredientes = recetaBD.Ingredientes
	}
	if recetaToUpdate.LitrosFinales == 0 {
		recetaToUpdate.LitrosFinales = recetaBD.LitrosFinales
	}

	return nil
}

func toEntities(ingredientesRequest []update_receta.Ingredientes) []entities.Ingredientes {
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

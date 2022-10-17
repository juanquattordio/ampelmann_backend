package update_producto_final

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Implementation struct {
	ProductoFinalProvider providers.ProductoFinal
}

var (
	ErrNotFound          = goErrors.New("product not found")
	ErrDuplicate         = goErrors.New("description already exists. Operation cancelled.")
	ErrAllreadyCancelled = goErrors.New("product's status is already 'disabled'. Operation cancelled.")
	ErrDisavailableUnit  = goErrors.New("unit of measurement disavailable. Operation cancelled.")
	ErrStatusRequired    = goErrors.New("status required is not available. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, id int64, request update_producto_final.RequestUpdate) (*entities.ProductoFinal, error) {

	// valida que exista la entidad a actualizar
	productDB, err := uc.ProductoFinalProvider.Search(&id, nil)
	if productDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	// si se quiere actualizar unidad de medida, valida su valor
	if request.Unidad != nil {
		if !isValidUnidad(*request.Unidad) {
			return nil, ErrDisavailableUnit
		}
	}

	if request.Descripcion == nil && request.Status == nil {
		productDB, err = changeStatusProductoFinal(productDB)
	} else {
		productDB, err = prepareToUpdate(uc, request, productDB)
	}
	if err != nil {
		return productDB, err
	}

	err = uc.ProductoFinalProvider.Update(productDB)
	if err != nil {
		return &entities.ProductoFinal{}, err
	}

	return productDB, nil
}
func changeStatusProductoFinal(productDB *entities.ProductoFinal) (*entities.ProductoFinal, error) {
	if productDB.Status != constants.Desactivo {
		productDB.Status = constants.Desactivo
		return productDB, nil
	} else {
		return productDB, ErrAllreadyCancelled
	}
}
func prepareToUpdate(uc *Implementation, request update_producto_final.RequestUpdate, productDB *entities.ProductoFinal) (*entities.ProductoFinal, error) {
	// si se quiere actualizar este campo, valida que no existan duplicados.
	if request.Descripcion != nil && productDB.Descripcion != *request.Descripcion {
		productExists, err := uc.ProductoFinalProvider.Search(nil, request.Descripcion)
		if productExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
			return nil, ErrDuplicate
		} else {
			productDB.Descripcion = *request.Descripcion
		}
	}
	if !isValidStatus(*request.Status) {
		return nil, ErrStatusRequired
	}
	// asigna los valores a actualizar, si corresponde
	if request.Status != nil && productDB.Status != strings.ToLower(*request.Status) {
		productDB.Status = strings.ToLower(*request.Status)
	}

	return productDB, nil
}

func isValidUnidad(unidad string) bool {
	unidad = strings.ToLower(unidad)
	return unidad == constants.LTS ||
		unidad == constants.UN
}

func isValidStatus(status string) bool {
	status = strings.ToLower(status)
	return status == constants.Activo ||
		status == constants.Desactivo
}

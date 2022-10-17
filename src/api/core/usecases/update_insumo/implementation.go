package update_insumo

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Implementation struct {
	InsumoProvider providers.Insumo
}

var (
	ErrNotFound          = goErrors.New("insumo not found")
	ErrDuplicate         = goErrors.New("name already exists. Operation cancelled.")
	ErrAllreadyCancelled = goErrors.New("insumo's status is already 'desactivo'. Operation cancelled.")
	ErrDisavailableUnit  = goErrors.New("unit of measurement disavailable. Operation cancelled.")
	ErrStatusRequired    = goErrors.New("status required is not available. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, id int64, request update_insumo.RequestUpdate) (*entities.Insumo, error) {
	// valida que exista la entidad a actualizar
	insumoDB, err := uc.InsumoProvider.Search(&id, nil)
	if insumoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	// si se quiere actualizar unidad de medida, valida su valor
	if request.Unidad != nil {
		if !isValidUnidad(*request.Unidad) {
			return nil, ErrDisavailableUnit
		}
	}

	if request.Nombre == nil && request.Status == nil && request.Stock == nil {
		insumoDB, err = changeStatusInsumo(insumoDB)
	} else {
		insumoDB, err = prepareToUpdate(uc, request, insumoDB)
	}
	if err != nil {
		return insumoDB, err
	}

	err = uc.InsumoProvider.Update(insumoDB)
	if err != nil {
		return &entities.Insumo{}, err
	}

	return insumoDB, nil
}
func changeStatusInsumo(insumoDB *entities.Insumo) (*entities.Insumo, error) {
	if insumoDB.Status != constants.Desactivo {
		insumoDB.Status = constants.Desactivo
		return insumoDB, nil
	} else {
		return insumoDB, ErrAllreadyCancelled
	}
}
func prepareToUpdate(uc *Implementation, request update_insumo.RequestUpdate, insumoDB *entities.Insumo) (*entities.Insumo, error) {
	// si se quiere actualizar este campo, valida que no existan duplicados.
	if request.Nombre != nil && insumoDB.Nombre != *request.Nombre {
		insumoExists, err := uc.InsumoProvider.Search(nil, request.Nombre)
		if insumoExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
			return nil, ErrDuplicate
		} else {
			insumoDB.Nombre = *request.Nombre
		}
	}

	if !isValidStatus(*request.Status) {
		return nil, ErrStatusRequired
	}
	// asigna los valores a actualizar, si corresponde
	if request.Status != nil && insumoDB.Status != strings.ToLower(*request.Status) {
		insumoDB.Status = strings.ToLower(*request.Status)
	}
	if request.Stock != nil && insumoDB.Stock != *request.Stock {
		insumoDB.Stock = *request.Stock
	}
	return insumoDB, nil
}

func isValidUnidad(unidad string) bool {
	unidad = strings.ToLower(unidad)
	return unidad == constants.KG ||
		unidad == constants.GRS ||
		unidad == constants.LTS ||
		unidad == constants.UN
}

func isValidStatus(status string) bool {
	status = strings.ToLower(status)
	return status == constants.Activo ||
		status == constants.Desactivo
}

package create_insumo

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Implementation struct {
	InsumoProvider providers.Insumo
}

var (
	ErrNotFound         = goErrors.New("insumo not found")
	ErrDuplicate        = goErrors.New("name already exists. Operation cancelled.")
	ErrDisavailableUnit = goErrors.New("unit of measurement disavailable. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_insumo.Request) (*entities.Insumo, error) {
	insumoExists, err := uc.InsumoProvider.Search(nil, request.Nombre)
	if insumoExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
	}

	if !isValidUnidad(*request.Unidad) {
		return nil, ErrDisavailableUnit
	}

	newInsumo := entities.NewInsumo(*request.Nombre, *request.Unidad, 0)

	err = uc.InsumoProvider.Save(*newInsumo)
	if err != nil {
		return &entities.Insumo{}, err
	}
	lastId, _ := uc.InsumoProvider.GetLastID()
	newInsumo.IdInsumo = lastId
	return newInsumo, nil
}

func isValidUnidad(unidad string) bool {
	unidad = strings.ToLower(unidad)
	return unidad == constants.KG ||
		unidad == constants.GRS ||
		unidad == constants.LTS ||
		unidad == constants.UN
}

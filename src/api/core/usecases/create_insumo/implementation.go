package create_insumo

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	InsumoProvider providers.Insumo
}

var (
	ErrNotFound    = goErrors.New("insumo not found")
	ErrDuplicate   = goErrors.New("name already exists. Operation cancelled.")
	ErrInternal    = goErrors.New("internal error")
	ErrWhCodeEmpty = goErrors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_insumo.Request) (*entities.Insumo, error) {
	insumoExists, err := uc.InsumoProvider.Search(nil, request.Nombre)
	if insumoExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, errors.NewInternalServer("Insumo ya existente")
	}

	newInsumo := entities.NewInsumo(*request.Nombre, *request.Stock)

	err = uc.InsumoProvider.Save(*newInsumo)
	if err != nil {
		return &entities.Insumo{}, err
	}
	lastId, _ := uc.InsumoProvider.GetLastID()
	newInsumo.IdInsumo = lastId
	return newInsumo, nil
}

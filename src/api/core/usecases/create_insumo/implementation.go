package create_insumo

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	InsumoProvider providers.Insumo
}

var (
	ErrNotFound  = goErrors.New("insumo not found")
	ErrDuplicate = goErrors.New("name already exists. Operation cancelled.")
	ErrInternal  = goErrors.New("internal error")
)

func (uc *Implementation) Execute(ctx context.Context, request create_insumo.Request) (*entities.Insumo, error) {
	insumoExists, err := uc.InsumoProvider.Search(nil, request.Nombre)
	if insumoExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
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

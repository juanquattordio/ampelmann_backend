package create_proveedor

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ProveedorProvider providers.Proveedor
}

var (
	ErrDuplicate   = goErrors.New("cuit already exists. Operation cancelled.")
	ErrInternal    = goErrors.New("internal error")
	ErrFieldsEmpty = goErrors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_proveedor.Request) (*entities.Proveedor, error) {
	proveedorExists, err := uc.ProveedorProvider.Search(nil, request.Cuit)
	if proveedorExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
	}

	newProveedor := entities.NewProveedor(*request.Cuit, *request.Nombre, *request.Ubicacion, *request.PaginaWeb)

	err = uc.ProveedorProvider.Save(*newProveedor)
	if err != nil {
		return &entities.Proveedor{}, err
	}
	lastId, _ := uc.ProveedorProvider.GetLastID()
	newProveedor.ID = lastId
	return newProveedor, nil
}

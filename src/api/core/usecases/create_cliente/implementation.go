package create_cliente

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ClienteProvider providers.Cliente
}

var (
	ErrDuplicate   = goErrors.New("cuit already exists. Operation cancelled.")
	ErrInternal    = goErrors.New("internal error")
	ErrFieldsEmpty = goErrors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_cliente.Request) (*entities.Cliente, error) {
	clienteExists, err := uc.ClienteProvider.Search(nil, request.Cuit)
	if clienteExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
	}

	newCliente := entities.NewCliente(*request.Cuit, *request.Nombre, *request.Ubicacion, *request.Email)

	err = uc.ClienteProvider.Save(*newCliente)
	if err != nil {
		return &entities.Cliente{}, err
	}
	lastId, _ := uc.ClienteProvider.GetLastID()
	newCliente.ID = lastId
	return newCliente, nil
}

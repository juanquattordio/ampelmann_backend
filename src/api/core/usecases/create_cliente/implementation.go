package create_cliente

import (
	"context"
	"errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ClienteProvider providers.Cliente
}

var (
	ErrNotFound    = errors.New("cliente not found")
	ErrDuplicate   = errors.New("cuit already exists. Operation cancelled.")
	ErrInternal    = errors.New("internal error")
	ErrWhCodeEmpty = errors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_cliente.Request) (*entities.Cliente, error) {

	newCliente := entities.NewCliente(*request.Cuit, *request.Nombre, *request.Ubicacion, *request.PaginaWeb)

	err := uc.ClienteProvider.Save(*newCliente)
	if err != nil {
		return &entities.Cliente{}, err
	}
	lastId, _ := uc.ClienteProvider.GetLastID()
	newCliente.ID = lastId
	return newCliente, nil
}

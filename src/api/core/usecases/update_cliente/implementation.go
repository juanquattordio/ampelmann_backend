package update_cliente

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	ClienteProvider providers.Cliente
}

var (
	ErrNotFound          = goErrors.New("cliente not found")
	ErrDuplicate         = goErrors.New("CUIT already exists. Operation cancelled")
	ErrAllreadyCancelled = goErrors.New("cliente's status is already 'desactivo'. Operation cancelled")
)

func (uc *Implementation) Execute(ctx context.Context, id int64, request update_cliente.Request) (*entities.Cliente, error) {

	// valida que exista la entidad a actualizar
	clienteDB, err := uc.ClienteProvider.Search(&id, nil)
	if clienteDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	if request.Cuit == nil && request.Nombre == nil && request.Ubicacion == nil && request.Email == nil && request.Status == nil {
		clienteDB, err = changeStatusCliente(clienteDB)
	} else {
		clienteDB, err = prepareToUpdate(uc, request, clienteDB)
	}
	if err != nil {
		return clienteDB, err
	}

	err = uc.ClienteProvider.Update(clienteDB)
	if err != nil {
		return &entities.Cliente{}, err
	}

	return clienteDB, nil
}
func changeStatusCliente(clienteDB *entities.Cliente) (*entities.Cliente, error) {
	if clienteDB.Status != constants.Desactivo {
		clienteDB.Status = constants.Desactivo
		return clienteDB, nil
	} else {
		return clienteDB, ErrAllreadyCancelled
	}
}
func prepareToUpdate(uc *Implementation, request update_cliente.Request, clienteDB *entities.Cliente) (*entities.Cliente, error) {
	// si se quiere actualizar este campo, valida que no existan duplicados.
	if request.Cuit != nil && clienteDB.Cuit != *request.Cuit {
		clienteExists, err := uc.ClienteProvider.Search(nil, request.Cuit)
		if clienteExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
			return nil, ErrDuplicate
		} else {
			clienteDB.Cuit = *request.Cuit
		}
	}

	// asigna los valores a actualizar, si corresponde
	if request.Nombre != nil && clienteDB.Nombre != *request.Nombre {
		clienteDB.Nombre = *request.Nombre
	}
	if request.Ubicacion != nil && clienteDB.Ubicacion != *request.Ubicacion {
		clienteDB.Ubicacion = *request.Ubicacion
	}
	if request.Email != nil && clienteDB.Email != *request.Email {
		clienteDB.Email = *request.Email
	}
	if request.Status != nil && clienteDB.Status != *request.Status {
		clienteDB.Status = *request.Status
	}

	return clienteDB, nil
}

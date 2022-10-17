package update_proveedor

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"strings"
)

type Implementation struct {
	ProveedorProvider providers.Proveedor
}

var (
	ErrNotFound          = goErrors.New("proveedor not found")
	ErrDuplicate         = goErrors.New("CUIT already exists. Operation cancelled")
	ErrAllreadyCancelled = goErrors.New("proveedor's status is already 'desactivo'. Operation cancelled")
	ErrStatusRequired    = goErrors.New("status required is not available. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, id int64, request update_proveedor.Request) (*entities.Proveedor, error) {

	// valida que exista la entidad a actualizar
	proveedorDB, err := uc.ProveedorProvider.Search(&id, nil)
	if proveedorDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	if request.Cuit == nil && request.Nombre == nil && request.Ubicacion == nil && request.PaginaWeb == nil && request.Status == nil {
		proveedorDB, err = changeStatusProveedor(proveedorDB)
	} else {
		proveedorDB, err = prepareToUpdate(uc, request, proveedorDB)
	}
	if err != nil {
		return proveedorDB, err
	}

	err = uc.ProveedorProvider.Update(proveedorDB)
	if err != nil {
		return &entities.Proveedor{}, err
	}

	return proveedorDB, nil
}
func changeStatusProveedor(proveedorDB *entities.Proveedor) (*entities.Proveedor, error) {
	if proveedorDB.Status != constants.Desactivo {
		proveedorDB.Status = constants.Desactivo
		return proveedorDB, nil
	} else {
		return proveedorDB, ErrAllreadyCancelled
	}
}
func prepareToUpdate(uc *Implementation, request update_proveedor.Request, proveedorDB *entities.Proveedor) (*entities.Proveedor, error) {
	// si se quiere actualizar este campo, valida que no existan duplicados.
	if request.Cuit != nil && proveedorDB.Cuit != *request.Cuit {
		proveedorExists, err := uc.ProveedorProvider.Search(nil, request.Cuit)
		if proveedorExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
			return nil, ErrDuplicate
		} else {
			proveedorDB.Cuit = *request.Cuit
		}
	}
	if !isValidStatus(*request.Status) {
		return nil, ErrStatusRequired
	}
	// asigna los valores a actualizar, si corresponde
	if request.Nombre != nil && proveedorDB.Nombre != *request.Nombre {
		proveedorDB.Nombre = *request.Nombre
	}
	if request.Ubicacion != nil && proveedorDB.Ubicacion != *request.Ubicacion {
		proveedorDB.Ubicacion = *request.Ubicacion
	}
	if request.PaginaWeb != nil && proveedorDB.PaginaWeb != *request.PaginaWeb {
		proveedorDB.PaginaWeb = *request.PaginaWeb
	}
	if request.Status != nil && proveedorDB.Status != strings.ToLower(*request.Status) {
		proveedorDB.Status = strings.ToLower(*request.Status)
	}

	return proveedorDB, nil
}

func isValidStatus(status string) bool {
	status = strings.ToLower(status)
	return status == constants.Activo ||
		status == constants.Desactivo
}

package update_deposito

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	DepositoProvider providers.Deposito
}

var (
	ErrNotFound          = goErrors.New("deposito not found")
	ErrDuplicate         = goErrors.New("deposito already exists. Operation cancelled")
	ErrAllreadyCancelled = goErrors.New("deposito's status is already 'desactivo'. Operation cancelled")
)

func (uc *Implementation) Execute(ctx context.Context, id int64, request update_deposito.Request) (*entities.Deposito, error) {

	// valida que exista la entidad a actualizar
	depositoDB, err := uc.DepositoProvider.Search(&id, nil)
	if depositoDB == nil && goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	if request.Descripcion == nil && request.Status == nil {
		depositoDB, err = changeStatusDeposito(depositoDB)
	} else {
		depositoDB, err = prepareToUpdate(uc, request, depositoDB)
	}
	if err != nil {
		return depositoDB, err
	}

	err = uc.DepositoProvider.Update(depositoDB)
	if err != nil {
		return &entities.Deposito{}, err
	}

	return depositoDB, nil
}
func changeStatusDeposito(depositoDB *entities.Deposito) (*entities.Deposito, error) {
	if depositoDB.Status != constants.Desactivo {
		depositoDB.Status = constants.Desactivo
		return depositoDB, nil
	} else {
		return depositoDB, ErrAllreadyCancelled
	}
}
func prepareToUpdate(uc *Implementation, request update_deposito.Request, depositoDB *entities.Deposito) (*entities.Deposito, error) {
	// si se quiere actualizar este campo, valida que no existan duplicados.
	if request.Descripcion != nil && depositoDB.Descripcion != *request.Descripcion {
		depositoExists, err := uc.DepositoProvider.Search(nil, request.Descripcion)
		if depositoExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
			return nil, ErrDuplicate
		} else {
			depositoDB.Descripcion = *request.Descripcion
		}
	}

	// asigna los valores a actualizar, si corresponde
	if request.Status != nil && depositoDB.Status != *request.Status {
		depositoDB.Status = *request.Status
	}

	return depositoDB, nil
}

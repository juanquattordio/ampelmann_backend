package create_deposito

import (
	"context"
	"database/sql"
	goErrors "errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
)

type Implementation struct {
	DepositoProvider providers.Deposito
}

var (
	ErrDuplicate   = goErrors.New("deposito already exists. Operation cancelled.")
	ErrInternal    = goErrors.New("internal error")
	ErrFieldsEmpty = goErrors.New("some fields can not be empty. Operation cancelled.")
)

func (uc *Implementation) Execute(ctx context.Context, request create_deposito.Request) (*entities.Deposito, error) {
	depositoExists, err := uc.DepositoProvider.Search(nil, request.Descripcion)
	if depositoExists != nil && !goErrors.Is(err, sql.ErrNoRows) {
		return nil, ErrDuplicate
	}

	newDeposito := entities.NewDeposito(*request.Descripcion)

	err = uc.DepositoProvider.Save(*newDeposito)
	if err != nil {
		return &entities.Deposito{}, err
	}
	lastId, _ := uc.DepositoProvider.GetLastID()
	newDeposito.ID = lastId
	return newDeposito, nil
}

package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type Deposito struct {
	ID          int64
	Descripcion string
	Status      string
}

func NewDeposito(descripcion string) *Deposito {
	deposito := &Deposito{
		Descripcion: descripcion,
		Status:      constants.Activo,
	}

	return deposito
}

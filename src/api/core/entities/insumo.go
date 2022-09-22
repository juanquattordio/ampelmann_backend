package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type Insumo struct {
	IdInsumo int64
	Nombre   string
	Stock    float64
	Status   string
}

func NewInsumo(descripcion string, stock float64) *Insumo {
	insumo := &Insumo{
		Nombre: descripcion,
		Stock:  stock,
		Status: constants.Activo,
	}

	return insumo
}

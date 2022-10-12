package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type ProductoFinal struct {
	Id          int64
	Descripcion string
	Status      string
}

func NewProductoFinal(descripcion string) *ProductoFinal {
	producto := &ProductoFinal{
		Descripcion: descripcion,
		Status:      constants.Activo,
	}

	return producto
}

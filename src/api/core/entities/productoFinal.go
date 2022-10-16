package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type ProductoFinal struct {
	Id          int64
	Descripcion string
	Stock       float64
	Status      string
}

func NewProductoFinal(descripcion string, stock float64) *ProductoFinal {
	producto := &ProductoFinal{
		Descripcion: descripcion,
		Stock:       stock,
		Status:      constants.Activo,
	}

	return producto
}

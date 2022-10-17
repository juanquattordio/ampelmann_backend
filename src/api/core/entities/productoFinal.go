package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type ProductoFinal struct {
	Id          int64
	Descripcion string
	Unidad      string
	Stock       float64
	Status      string
}

func NewProductoFinal(descripcion string, unidad string, stock float64) *ProductoFinal {
	producto := &ProductoFinal{
		Descripcion: descripcion,
		Unidad:      unidad,
		Stock:       stock,
		Status:      constants.Activo,
	}

	return producto
}

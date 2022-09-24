package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type Proveedor struct {
	ID        int64
	Cuit      string
	Nombre    string
	Ubicacion string
	PaginaWeb string
	Status    string
}

func NewProveedor(cuit string, nombre string, ubicacion string, paginaWeb string) *Proveedor {
	proveedor := &Proveedor{
		Cuit:      cuit,
		Nombre:    nombre,
		Ubicacion: ubicacion,
		PaginaWeb: paginaWeb,
		Status:    constants.Activo,
	}

	return proveedor
}

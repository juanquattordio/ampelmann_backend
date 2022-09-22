package entities

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"

type Cliente struct {
	ID        int64
	Cuit      string
	Nombre    string
	Ubicacion string
	Email     string
	Status    string
}

func NewCliente(cuit string, nombre string, ubicacion string, email string) *Cliente {
	cliente := &Cliente{
		Cuit:      cuit,
		Nombre:    nombre,
		Ubicacion: ubicacion,
		Email:     email,
		Status:    constants.Activo,
	}

	return cliente
}

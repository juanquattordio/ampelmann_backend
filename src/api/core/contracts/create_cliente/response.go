package create_cliente

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID        string `json:"cliente_id"`
	Cuit      string `json:"cuit"`
	Nombre    string `json:"cliente_nombre"`
	Ubicacion string `json:"cliente_ubicacion"`
	PaginaWeb string `json:"cliente_web"`
}

func NewResponse(cliente *entities.Cliente) *Response {
	return &Response{
		ID:        fmt.Sprintf("%d", cliente.ID),
		Cuit:      cliente.Cuit,
		Nombre:    cliente.Nombre,
		Ubicacion: cliente.Ubicacion,
		PaginaWeb: cliente.PaginaWeb,
	}
}

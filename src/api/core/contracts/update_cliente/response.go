package update_cliente

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID        string `json:"cliente_id"`
	Cuit      string `json:"cuit"`
	Nombre    string `json:"cliente_nombre"`
	Ubicacion string `json:"cliente_ubicacion"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}

func NewResponse(cliente *entities.Cliente) *Response {
	return &Response{
		ID:        fmt.Sprintf("%d", cliente.ID),
		Cuit:      cliente.Cuit,
		Nombre:    cliente.Nombre,
		Ubicacion: cliente.Ubicacion,
		Email:     cliente.Email,
		Status:    cliente.Status,
	}
}

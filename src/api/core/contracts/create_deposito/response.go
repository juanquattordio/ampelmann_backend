package create_deposito

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID          string `json:"deposito_id"`
	Descripcion string `json:"descripcion"`
	Status      string `json:"status"`
}

func NewResponse(deposito *entities.Deposito) *Response {
	return &Response{
		ID:          fmt.Sprintf("%d", deposito.ID),
		Descripcion: deposito.Descripcion,
		Status:      deposito.Status,
	}
}

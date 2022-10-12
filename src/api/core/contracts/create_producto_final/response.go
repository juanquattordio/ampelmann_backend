package create_producto_final

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID          string `json:"id_producto_final"`
	Descripcion string `json:"descripcion"`
	Status      string `json:"status"`
}

func NewResponse(producto_final *entities.ProductoFinal) *Response {
	return &Response{
		ID:          fmt.Sprintf("%d", producto_final.Id),
		Descripcion: producto_final.Descripcion,
		Status:      producto_final.Status,
	}
}

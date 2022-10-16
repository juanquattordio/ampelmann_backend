package update_producto_final

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID          string `json:"id_producto"`
	Descripcion string `json:"descripcion"`
	Status      string `json:"status"`
}

func NewResponse(product *entities.ProductoFinal) *Response {
	return &Response{
		ID:          fmt.Sprintf("%d", product.Id),
		Descripcion: product.Descripcion,
		Status:      product.Status,
	}
}

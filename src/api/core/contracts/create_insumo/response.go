package create_insumo

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID     string `json:"insumo_id"`
	Nombre string `json:"insumo_nombre"`
	Stock  string `json:"insumo_stock"`
	Status string `json:"insumo_status"`
}

func NewResponse(insumo *entities.Insumo) *Response {
	return &Response{
		ID:     fmt.Sprintf("%d", insumo.IdInsumo),
		Nombre: insumo.Nombre,
		Stock:  fmt.Sprintf("%f", insumo.Stock),
		Status: insumo.Status,
	}
}

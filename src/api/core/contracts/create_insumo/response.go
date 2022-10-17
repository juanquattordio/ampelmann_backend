package create_insumo

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID     string `json:"insumo_id"`
	Nombre string `json:"insumo_nombre"`
	Status string `json:"insumo_status"`
}

func NewResponse(insumo *entities.Insumo) *Response {
	return &Response{
		ID:     fmt.Sprintf("%d", insumo.IdInsumo),
		Nombre: insumo.Nombre,
		Status: insumo.Status,
	}
}

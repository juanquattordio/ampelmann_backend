package search_proveedor

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	ID        string `json:"proveedor_id"`
	Cuit      string `json:"cuit"`
	Nombre    string `json:"proveedor_nombre"`
	Ubicacion string `json:"proveedor_ubicacion"`
	PaginaWeb string `json:"pagina_web"`
	Status    string `json:"status"`
}

func NewResponse(proveedor *entities.Proveedor) *Response {
	return &Response{
		ID:        fmt.Sprintf("%d", proveedor.ID),
		Cuit:      proveedor.Cuit,
		Nombre:    proveedor.Nombre,
		Ubicacion: proveedor.Ubicacion,
		PaginaWeb: proveedor.PaginaWeb,
		Status:    proveedor.Status,
	}
}

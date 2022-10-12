package receta

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	IdHeader        *int64   `form:"id_header" json:"id_header"`
	DetallePasoPaso *string  `form:"detalle_paso_paso" json:"detalle_paso_paso"`
	IdProductoFinal *int64   `form:"id_producto_final" json:"id_producto_final"`
	LitrosFinales   *float64 `form:"litros_finales" json:"litros_finales"`
	Insumos         []Insumo `form:"insumos" json:"insumos"`
}

func NewResponse(receta *entities.RecetaHeader) *Response {
	insumos := make([]Insumo, len(receta.LineasInsumos))
	for i := range receta.LineasInsumos {
		insumos[i].IdInsumo = &receta.LineasInsumos[i].IdInsumo
		insumos[i].UnidadMedida = &receta.LineasInsumos[i].UnidadMedida
		insumos[i].Cantidad = &receta.LineasInsumos[i].Cantidad

	}
	return &Response{
		IdHeader:        &receta.IdHeader,
		DetallePasoPaso: &receta.PasoPaso,
		IdProductoFinal: receta.IdProductoFinal,
		LitrosFinales:   &receta.LitrosFinales,
		Insumos:         insumos,
	}
}

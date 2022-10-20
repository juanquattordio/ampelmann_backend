package update_receta

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type Response struct {
	IdHeader        *int64         `form:"id_header" json:"id_header"`
	DetallePasoPaso *string        `form:"detalle_paso_paso" json:"detalle_paso_paso"`
	IdProductoFinal *int64         `form:"id_producto" json:"id_producto"`
	LitrosFinales   *float64       `form:"litros_finales" json:"litros_finales"`
	Ingredientes    []Ingredientes `form:"ingredientes" json:"ingredientes"`
}

func NewResponse(receta *entities.RecetaHeader) *Response {
	ingredientes := make([]Ingredientes, len(receta.Ingredientes))
	for i := range receta.Ingredientes {
		ingredientes[i].IdInsumo = &receta.Ingredientes[i].IdInsumo
		ingredientes[i].UnidadMedida = receta.Ingredientes[i].UnidadMedida
		ingredientes[i].Cantidad = &receta.Ingredientes[i].Cantidad

	}
	return &Response{
		IdHeader:        &receta.IdHeader,
		DetallePasoPaso: &receta.PasoPaso,
		IdProductoFinal: receta.IdProductoFinal,
		LitrosFinales:   &receta.LitrosFinales,
		Ingredientes:    ingredientes,
	}
}

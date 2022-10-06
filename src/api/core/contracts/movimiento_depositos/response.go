package movimiento_depositos

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

type Response struct {
	Fecha             time.Time `form:"fecha" json:"fecha"`
	IdDepositoOrigen  *int64    `form:"id_deposito_origen" json:"id_deposito_origen"`
	IdDepositoDestino *int64    `form:"id_deposito_destino" json:"id_deposito_destino"`
	Insumos           []Insumo  `form:"insumos" json:"insumos"`
	Status            *string   `form:"status" json:"status"`
}

func NewResponse(movimiento *entities.MovimientoHeader) *Response {
	insumos := make([]Insumo, len(movimiento.Lineas))
	for i := range movimiento.Lineas {
		insumos[i].IdInsumo = &movimiento.Lineas[i].IdInsumo
		insumos[i].Cantidad = &movimiento.Lineas[i].Cantidad
		insumos[i].Obseraciones = movimiento.Lineas[i].Obseraciones
	}
	return &Response{
		Fecha:             movimiento.Fecha,
		IdDepositoOrigen:  &movimiento.IdDepositoOrigen,
		IdDepositoDestino: &movimiento.IdDepositoDestino,
		Insumos:           insumos,
		Status:            &movimiento.Status,
	}
}

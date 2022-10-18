package movimiento_depositos

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"time"
)

type Response struct {
	IdHeader          *int64    `form:"id_header" json:"id_header"`
	Fecha             time.Time `form:"fecha" json:"fecha"`
	IdDepositoOrigen  *int64    `form:"id_deposito_origen" json:"id_deposito_origen"`
	IdDepositoDestino *int64    `form:"id_deposito_destino" json:"id_deposito_destino"`
	CausaMovimiento   string    `form:"causa_movimiento" json:"causa_movimiento"`
	Insumos           []Insumo  `form:"insumos" json:"insumos"`
	Status            *string   `form:"status" json:"status"`
}

func NewResponse(movimiento *entities.MovimientoHeader) *Response {
	insumos := make([]Insumo, len(movimiento.Lineas))
	for i := range movimiento.Lineas {
		insumos[i].IdLinea = &movimiento.Lineas[i].IdLinea
		insumos[i].IdInsumo = &movimiento.Lineas[i].IdInsumo
		insumos[i].Cantidad = &movimiento.Lineas[i].Cantidad
		insumos[i].Observaciones = movimiento.Lineas[i].Observaciones
	}
	return &Response{
		IdHeader:          &movimiento.IdHeader,
		Fecha:             movimiento.Fecha,
		IdDepositoOrigen:  &movimiento.IdDepositoOrigen,
		IdDepositoDestino: &movimiento.IdDepositoDestino,
		CausaMovimiento:   movimiento.CausaMovimiento,
		Insumos:           insumos,
		Status:            &movimiento.Status,
	}
}

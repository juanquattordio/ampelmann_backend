package entities

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"time"
)

type MovimientoProductoHeader struct {
	IdHeader          int64
	Fecha             time.Time
	IdDepositoOrigen  int64
	IdDepositoDestino int64
	Lineas            []MovimientoProductoLine
	Status            string
	CausaMovimiento   string
}

type MovimientoProductoLine struct {
	IdLinea       int64
	Id            int64
	Cantidad      float64
	Observaciones string
}

func NewMovimientoProductoDeposito(origen int64, destino int64, lineas []MovimientoProductoLine, causaMovimiento string) *MovimientoProductoHeader {
	movimientoHeader := &MovimientoProductoHeader{
		Fecha:             time.Now().UTC(),
		IdDepositoOrigen:  origen,
		IdDepositoDestino: destino,
		Lineas:            lineas,
		Status:            constants.Activo,
		CausaMovimiento:   causaMovimiento,
	}

	return movimientoHeader
}

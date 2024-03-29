package entities

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"time"
)

type MovimientoHeader struct {
	IdHeader          int64
	Fecha             time.Time
	IdDepositoOrigen  int64
	IdDepositoDestino int64
	Lineas            []MovimientoLine
	Status            string
	CausaMovimiento   string
}

type MovimientoLine struct {
	IdLinea       int64
	IdArticulo    int64
	Cantidad      float64
	Observaciones string
}

func NewMovimientoDeposito(origen int64, destino int64, lineas []MovimientoLine, causaMovimiento string) *MovimientoHeader {
	movimientoHeader := &MovimientoHeader{
		Fecha:             time.Now().UTC(),
		IdDepositoOrigen:  origen,
		IdDepositoDestino: destino,
		Lineas:            lineas,
		Status:            constants.Activo,
		CausaMovimiento:   causaMovimiento,
	}

	return movimientoHeader
}

package entities

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities/constants"
	"time"
)

type MovimientoHeader struct {
	Fecha             time.Time
	IdDepositoOrigen  int64
	IdDepositoDestino int64
	Lineas            []MovimientoLine
	Status            string
}

type MovimientoLine struct {
	IdInsumo     int64
	Cantidad     float64
	Obseraciones string
}

func NewMovimientoDeposito(origen int64, destino int64, lineas []MovimientoLine) *MovimientoHeader {
	movimientoHeader := &MovimientoHeader{
		Fecha:             time.Now().UTC(),
		IdDepositoOrigen:  origen,
		IdDepositoDestino: destino,
		Lineas:            lineas,
		Status:            constants.Activo,
	}

	return movimientoHeader
}

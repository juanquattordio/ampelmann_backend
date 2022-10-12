package entities

type RecetaHeader struct {
	IdHeader        int64
	PasoPaso        string
	IdProductoFinal *int64
	LineasInsumos   []RecetaInsumo
	LitrosFinales   float64
}

type RecetaInsumo struct {
	IdLinea      int64
	IdInsumo     int64
	UnidadMedida string
	Cantidad     float64
}

func NewRecetaDeposito(pasoPaso string, idProductoFinal *int64, lineas []RecetaInsumo, litrosFinales float64) *RecetaHeader {
	recetaHeader := &RecetaHeader{
		PasoPaso:        pasoPaso,
		IdProductoFinal: idProductoFinal,
		LineasInsumos:   lineas,
		LitrosFinales:   litrosFinales,
	}

	return recetaHeader
}

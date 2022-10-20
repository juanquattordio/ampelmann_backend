package entities

type RecetaHeader struct {
	IdHeader        int64
	PasoPaso        string
	IdProductoFinal *int64
	Ingredientes    []Ingredientes
	LitrosFinales   float64
}

type Ingredientes struct {
	IdInsumo      int64
	UnidadMedida  string
	Cantidad      float64
	Observaciones string
}

func NewReceta(idHeader int64, pasoPaso string, idProductoFinal *int64, lineas []Ingredientes, litrosFinales float64) *RecetaHeader {
	recetaHeader := &RecetaHeader{
		IdHeader:        idHeader,
		PasoPaso:        pasoPaso,
		IdProductoFinal: idProductoFinal,
		Ingredientes:    lineas,
		LitrosFinales:   litrosFinales,
	}

	return recetaHeader
}

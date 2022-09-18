package entities

type Insumo struct {
	IdInsumo int64
	Nombre   string
	Stock    float64
	Status   string
}

func NewInsumo(descripcion string, stock float64) *Insumo {
	insumo := &Insumo{
		Nombre: descripcion,
		Stock:  stock,
		Status: "activo",
	}

	return insumo
}

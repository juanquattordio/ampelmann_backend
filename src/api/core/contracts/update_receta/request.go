package update_receta

type Request struct {
	DetallePasoPaso *string        `form:"detalle_paso_paso" json:"detalle_paso_paso"`
	IdProductoFinal *int64         `form:"id_producto" json:"id_producto"`
	Ingredientes    []Ingredientes `form:"ingredientes" json:"ingredientes"`
	LitrosFinales   *float64       `form:"litros_finales" json:"litros_finales"`
}

type Ingredientes struct {
	IdInsumo      *int64   `form:"id_insumo" json:"id_insumo"`
	Cantidad      *float64 `form:"cantidad" json:"cantidad"`
	UnidadMedida  string   `form:"unidad_medida" json:"unidad_medida"`
	Observaciones string   `form:"observaciones" json:"observaciones"`
}

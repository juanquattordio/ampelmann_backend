package receta

type Request struct {
	DetallePasoPaso *string  `form:"detalle_paso_paso" json:"detalle_paso_paso" binding:"required"`
	IdProductoFinal *int64   `form:"id_producto" json:"id_producto"`
	Insumos         []Insumo `form:"insumos" json:"insumos" binding:"required"`
	LitrosFinales   *float64 `form:"litros_finales" json:"litros_finales" binding:"required"`
}

type Insumo struct {
	IdLinea      *int64   `form:"id_linea" json:"id_linea"`
	IdInsumo     *int64   `form:"id_insumo" json:"id_insumo" binding:"required"`
	UnidadMedida *string  `form:"unidad_medida" json:"unidad_medida" binding:"required"`
	Cantidad     *float64 `form:"cantidad" json:"cantidad" binding:"required"`
}

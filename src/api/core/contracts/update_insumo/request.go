package update_insumo

type RequestUpdate struct {
	Nombre *string  `form:"nombre" json:"nombre"`
	Unidad *string  `form:"unidad" json:"unidad"`
	Stock  *float64 `form:"stock" json:"stock"`
	Status *string  `form:"status" json:"status"`
}

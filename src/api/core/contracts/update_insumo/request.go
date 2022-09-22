package update_insumo

type RequestUpdate struct {
	Nombre *string  `form:"nombre" json:"nombre"`
	Stock  *float64 `form:"stock" json:"stock"`
	Status *string  `form:"status" json:"status"`
}

type RequestDelete struct {
	Id *int `form:"insumo_id" json:"insumo_id"`
}

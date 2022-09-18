package create_insumo

type Request struct {
	Nombre *string  `form:"nombre" json:"nombre" binding:"required"`
	Stock  *float64 `form:"stock" json:"stock"`
}

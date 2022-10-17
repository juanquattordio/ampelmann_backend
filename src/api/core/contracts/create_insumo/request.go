package create_insumo

type Request struct {
	Nombre *string `form:"nombre" json:"nombre" binding:"required"`
	Unidad *string `form:"unidad" json:"unidad" binding:"required"`
	Status *string `form:"status" json:"status"`
}

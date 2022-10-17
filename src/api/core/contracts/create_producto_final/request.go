package create_producto_final

type Request struct {
	Descripcion *string `form:"descripcion" json:"descripcion" binding:"required"`
	Unidad      *string `form:"unidad" json:"unidad" binding:"required"`
	Status      *string `form:"status" json:"status"`
}

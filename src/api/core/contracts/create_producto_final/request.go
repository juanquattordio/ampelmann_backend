package create_producto_final

type Request struct {
	Descripcion *string  `form:"descripcion" json:"descripcion" binding:"required"`
	Stock       *float64 `form:"stock" json:"stock"`
	Status      *string  `form:"status" json:"status"`
}

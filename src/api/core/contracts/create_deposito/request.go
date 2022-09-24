package create_deposito

type Request struct {
	Descripcion *string `form:"descripcion" json:"descripcion" binding:"required"`
	Status      *string `form:"status" json:"status" binding:""`
}

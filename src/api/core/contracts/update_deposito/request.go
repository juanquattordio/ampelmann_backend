package update_deposito

type Request struct {
	Descripcion *string `form:"descripcion" json:"descripcion"`
	Status      *string `form:"status" json:"status" binding:""`
}

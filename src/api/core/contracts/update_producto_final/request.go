package update_producto_final

type RequestUpdate struct {
	Descripcion *string `form:"descripcion" json:"descripcion"`
	Status      *string `form:"status" json:"status"`
}

type RequestDelete struct {
	Id *int `form:"insumo_id" json:"id_producto"`
}

package update_producto_final

type RequestUpdate struct {
	Descripcion *string `form:"descripcion" json:"descripcion"`
	Unidad      *string `form:"unidad" json:"unidad"`
	Status      *string `form:"status" json:"status"`
}

type RequestDelete struct {
	Id *int `form:"insumo_id" json:"id_producto"`
}

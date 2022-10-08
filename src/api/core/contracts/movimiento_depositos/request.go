package movimiento_depositos

type Request struct {
	IdDepositoOrigen  *int64   `form:"id_deposito_origen" json:"id_deposito_origen" binding:"required"`
	IdDepositoDestino *int64   `form:"id_deposito_destino" json:"id_deposito_destino" binding:"required"`
	Insumos           []Insumo `form:"insumos" json:"insumos" binding:"required"`
	Status            *string  `form:"status" json:"status" binding:""`
}

type Insumo struct {
	IdLinea      *int64   `form:"id_linea" json:"id_linea"`
	IdInsumo     *int64   `form:"id_insumo" json:"id_insumo" binding:"required"`
	Cantidad     *float64 `form:"cantidad" json:"cantidad" binding:"required"`
	Obseraciones string   `form:"observaciones" json:"observaciones"`
}

package movimiento_depositos

type Request struct {
	IdDepositoOrigen  int64       `form:"id_deposito_origen" json:"id_deposito_origen"`
	IdDepositoDestino int64       `form:"id_deposito_destino" json:"id_deposito_destino"`
	Insumos           []Articulos `form:"insumos" json:"insumos,omitempty"`
	Productos         []Articulos `form:"productos" json:"productos,omitempty"`
	Status            *string     `form:"status" json:"status" binding:""`
	CausaMovimiento   string      `form:"causa_movimiento" json:"causa_movimiento" binding:""`
}

type Articulos struct {
	IdLinea       int64    `form:"id_linea" json:"id_linea"`
	IdArticulo    *int64   `form:"id" json:"id" binding:"required"`
	Cantidad      *float64 `form:"cantidad" json:"cantidad" binding:"required"`
	Observaciones string   `form:"observaciones" json:"observaciones"`
}

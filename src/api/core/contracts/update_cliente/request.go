package update_cliente

type Request struct {
	Cuit      *string `form:"cuit" json:"cuit"`
	Nombre    *string `form:"nombre" json:"nombre"`
	Ubicacion *string `form:"ubicacion" json:"ubicacion"`
	Email     *string `form:"email" json:"email" binding:""`
	Status    *string `form:"status" json:"status" binding:""`
}

package create_proveedor

type Request struct {
	Cuit      *string `form:"cuit" json:"cuit" binding:"required"`
	Nombre    *string `form:"nombre" json:"nombre" binding:"required"`
	Ubicacion *string `form:"ubicacion" json:"ubicacion" binding:"required"`
	PaginaWeb *string `form:"pagina_web" json:"pagina_web" binding:""`
	Status    *string `form:"status" json:"status" binding:""`
}

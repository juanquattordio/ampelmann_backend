package update_proveedor

type Request struct {
	Cuit      *string `form:"cuit" json:"cuit"`
	Nombre    *string `form:"nombre" json:"nombre"`
	Ubicacion *string `form:"ubicacion" json:"ubicacion"`
	PaginaWeb *string `form:"pagina_web" json:"pagina_web" binding:""`
	Status    *string `form:"status" json:"status" binding:""`
}

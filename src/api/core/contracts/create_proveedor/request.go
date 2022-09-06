package create_cliente

type Request struct {
	Cuit      *string `form:"cuit" json:"cuit" binding:"required"`
	Nombre    *string `form:"nombre" json:"nombre" binding:"required"`
	Ubicacion *string `form:"ubicacion" json:"ubicacion" binding:"required"`
	Email     *string `form:"email" json:"email" binding:""`
	Status    *string `form:"status" json:"status" binding:""`
}

func (command Request) Validate() error {
	return nil
}

package update_proveedor

import "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura_compra"

type Request struct {
	Cuit      *string `form:"cuit" json:"cuit"`
	Nombre    *string `form:"nombre" json:"nombre"`
	Ubicacion *string `form:"ubicacion" json:"ubicacion"`
	PaginaWeb *string `form:"pagina_web" json:"pagina_web" binding:""`
	Status    *string `form:"status" json:"status" binding:""`
}

type RequestUpdateHistorialPrecio struct {
	IdProveedor    *int64                                  `form:"id_proveedor" json:"id_proveedor" binding:"required"`
	IdInsumo       *int64                                  `form:"id_insumo" json:"id_insumo" binding:"required"`
	PrecioUnitario *float64                                `form:"precio_unitario" json:"precio_unitario"`
	Fecha          create_factura_compra.CustomFechaOrigen `form:"fecha" json:"fecha"`
	Status         string                                  `form:"status" json:"status"`
}

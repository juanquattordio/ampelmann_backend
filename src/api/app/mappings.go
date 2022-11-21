package app

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/infrastructure/dependencies"
)

func configureMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	configureAPIMappings(router, handlers)
}

func configureAPIMappings(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	ampelmannGroup := router.Group("/ampelmann")

	// Clientes endpoints
	clientes := ampelmannGroup.Group("/clientes")
	clientes.POST("", handlers.CreateCliente.Handle)
	clientes.GET("", handlers.SearchCliente.Handle)
	clientes.PATCH("/:id", handlers.UpdateCliente.Handle)
	clientes.PATCH("/:id/cancel", handlers.UpdateCliente.Handle)

	// Insumos endpoints
	insumos := ampelmannGroup.Group("/insumos")
	insumos.POST("", handlers.CreateInsumo.Handle)
	insumos.GET("", handlers.SearchInsumo.Handle)
	insumos.PATCH("/:id", handlers.UpdateInsumo.Handle)
	insumos.PATCH("/:id/cancel", handlers.UpdateInsumo.Handle)
	insumos.GET("/stock", handlers.GetStockInsumo.Handle)

	// Productos Finales endpoints
	productosFinales := ampelmannGroup.Group("/productos")
	productosFinales.POST("", handlers.CreateProductoFinal.Handle)
	productosFinales.GET("", handlers.SearchProductoFinal.Handle)
	productosFinales.PATCH("/:id", handlers.UpdateProductoFinal.Handle)
	productosFinales.PATCH("/:id/cancel", handlers.UpdateProductoFinal.Handle)
	productosFinales.GET("/stock", handlers.GetStockProductoFinal.Handle)

	// Proveedores endpoints
	proveedores := ampelmannGroup.Group("/proveedores")
	proveedores.POST("", handlers.CreateProveedor.Handle)
	proveedores.GET("", handlers.SearchProveedor.Handle)
	proveedores.PATCH("/:id", handlers.UpdateProveedor.Handle)
	proveedores.PATCH("/:id/cancel", handlers.UpdateProveedor.Handle)
	proveedores.PATCH("/historial-precios", handlers.UpdateHistorial.Handle)

	// Depositos endpoints
	depositos := ampelmannGroup.Group("/depositos")
	depositos.POST("", handlers.CreateDeposito.Handle)
	depositos.PATCH("/:id", handlers.UpdateDeposito.Handle)
	depositos.PATCH("/:id/cancel", handlers.UpdateDeposito.Handle)
	depositos.POST("movimiento/create", handlers.CreateMovimientoDeposito.Handle)

	// Facturas
	facturas := ampelmannGroup.Group("/facturas")
	facturas.POST("compra/create", handlers.CreateFacturaCompra.Handle)
	facturas.POST("venta/create", handlers.CreateFacturaVenta.Handle)

	// Recetas
	recetas := ampelmannGroup.Group("/recetas")
	recetas.POST("", handlers.CreateReceta.Handle)
	recetas.PATCH("/:id", handlers.UpdateReceta.Handle)
	recetas.DELETE("/:id", handlers.DeleteReceta.Handle)

	// Batchs Produccion
	batchs := ampelmannGroup.Group("/batch")
	batchs.POST("", handlers.CreateBatch.Handle)

	// Reportes Informes
	reports := ampelmannGroup.Group("/reports")
	reports.GET("/insumos/stock-desactivados", handlers.InsumosReports.Handle)
	reports.GET("/productos/stock-desactivados", handlers.ProductosReports.Handle)
	reports.GET("/clientes/desactivados", handlers.ClientesReports.Handle)
	reports.GET("/facturacion-ventas", handlers.VentasReports.Handle)
}

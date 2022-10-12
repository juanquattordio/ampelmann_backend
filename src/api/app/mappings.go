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
	insumos.POST("movimiento/create", handlers.CreateMovimientoDeposito.Handle)

	// Productos Finales endpoints
	productosFinales := ampelmannGroup.Group("/producto_final")
	productosFinales.POST("", handlers.CreateProductoFinal.Handle)
	productosFinales.GET("", handlers.SearchProductoFinal.Handle)
	productosFinales.PATCH("/:id", handlers.UpdateProductoFinal.Handle)
	productosFinales.PATCH("/:id/cancel", handlers.UpdateProductoFinal.Handle)
	//productosFinales.GET("/stock", handlers.GetStockInsumo.Handle)
	//productosFinales.POST("movimiento/create", handlers.CreateMovimientoDeposito.Handle)

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

	// Facturas
	facturasCompra := ampelmannGroup.Group("/facturas-compra")
	facturasCompra.POST("/create", handlers.CreateFacturaCompra.Handle)
}

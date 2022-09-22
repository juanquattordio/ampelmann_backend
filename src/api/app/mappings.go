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

	// Insumos endpoints
	insumos := ampelmannGroup.Group("/insumos")
	insumos.POST("", handlers.CreateInsumo.Handle)
	insumos.GET("", handlers.SearchInsumo.Handle)
	insumos.PATCH("/:id", handlers.UpdateInsumo.Handle)

}

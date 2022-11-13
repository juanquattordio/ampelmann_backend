package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_factura_venta"
	"net/http"
)

type CreateFacturaVenta struct {
	CreateFacturaVentaUseCase create_factura_venta.UseCase
}

func (handler CreateFacturaVenta) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateFacturaVenta) handle(ctx *gin.Context) {
	var request contracts.RequestFacturaVenta

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}
	factura, err := handler.CreateFacturaVentaUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		// todo mejorar estos errores
		case create_cliente.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	facturaCompra, err := factura.ToFacturaCompra()
	ctx.JSON(http.StatusCreated, contracts.NewResponse(facturaCompra))

}

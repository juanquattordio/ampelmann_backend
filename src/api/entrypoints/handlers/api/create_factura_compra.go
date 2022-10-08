package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_factura_compra"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_factura_compra"
	"net/http"
)

type CreateFacturaCompra struct {
	CreateFacturaCompraUseCase create_factura_compra.UseCase
}

func (handler CreateFacturaCompra) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateFacturaCompra) handle(ctx *gin.Context) {
	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	factura, err := handler.CreateFacturaCompraUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		// todo mejorar estos errores
		case create_cliente.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(factura))

}

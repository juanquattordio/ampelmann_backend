package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_producto_final"
	"net/http"
)

type CreateProductoFinal struct {
	CreateProductoFinalUseCase create_producto_final.UseCase
}

func (handler CreateProductoFinal) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateProductoFinal) handle(ctx *gin.Context) {

	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	newProductoFinal, err := handler.CreateProductoFinalUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		case create_insumo.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(newProductoFinal))

}

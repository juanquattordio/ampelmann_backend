package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/movimiento_depositos"
	"net/http"
)

type CreateMovimientoDeposito struct {
	CreateMovimientoDepositoUseCase movimiento_depositos.UseCase
}

func (handler CreateMovimientoDeposito) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateMovimientoDeposito) handle(ctx *gin.Context) {
	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	if request.Insumos != nil && request.Productos != nil {
		ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Only one type of articles: insumos or productos"))
		return
	}

	movimiento, err := handler.CreateMovimientoDepositoUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		// todo mejorar estos errores
		case create_cliente.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(movimiento))

}

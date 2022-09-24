package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_deposito"
	"net/http"
)

type CreateDeposito struct {
	CreateDepositoUseCase create_deposito.UseCase
}

func (handler CreateDeposito) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateDeposito) handle(ctx *gin.Context) {
	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	newDeposito, err := handler.CreateDepositoUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		case create_deposito.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(newDeposito))

}

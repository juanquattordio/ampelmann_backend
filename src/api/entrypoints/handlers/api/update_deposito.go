package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_deposito"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_deposito"
	"net/http"
	"strconv"
)

type UpdateDeposito struct {
	UpdateDepositoUseCase update_deposito.UseCase
}

func (handler UpdateDeposito) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler UpdateDeposito) handle(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, web.NewResponse(http.StatusBadRequest, nil, "Parametro incorrecto"))
		return
	}

	var request contracts.Request

	if err = ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		return
	}

	depositoUpdated, err := handler.UpdateDepositoUseCase.Execute(ctx, id, request)

	if err != nil {
		switch err {
		case update_deposito.ErrDuplicate:
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		case update_deposito.ErrNotFound:
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		case update_deposito.ErrAllreadyCancelled:
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	} else {
		ctx.JSON(http.StatusOK, contracts.NewResponse(depositoUpdated))
	}

}

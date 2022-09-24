package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_cliente"
	"net/http"
	"strconv"
)

type UpdateCliente struct {
	UpdateClienteUseCase update_cliente.UseCase
}

func (handler UpdateCliente) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler UpdateCliente) handle(ctx *gin.Context) {

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

	clienteUpdated, err := handler.UpdateClienteUseCase.Execute(ctx, id, request)

	if err != nil {
		switch err {
		case update_cliente.ErrDuplicate:
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		case update_cliente.ErrNotFound:
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		case update_cliente.ErrAllreadyCancelled:
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	} else {
		ctx.JSON(http.StatusOK, contracts.NewResponse(clienteUpdated))
	}

}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"net/http"
)

type CreateCliente struct {
	CreateClienteUseCase create_cliente.UseCase
}

func (handler CreateCliente) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateCliente) handle(ctx *gin.Context) {

	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	newCliente, err := handler.CreateClienteUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		case create_cliente.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(newCliente))

}

package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_cliente"
	"net/http"
	"strconv"
)

type SearchCliente struct {
	SearchClienteUseCase search_cliente.UseCase
}

func (handler SearchCliente) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler SearchCliente) handle(ctx *gin.Context) {
	var request contracts.Request
	id := ctx.Query("id")
	if id != "" {
		clienteId, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.Id = &clienteId
	}
	cuit := ctx.Query("cuit")
	request.Cuit = &cuit

	clienteResult, err := handler.SearchClienteUseCase.Execute(ctx, request.Id, request.Cuit)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, "Source not found"))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponse(clienteResult))

}

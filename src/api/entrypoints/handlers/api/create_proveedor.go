package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_proveedor"
	"net/http"
)

type CreateProveedor struct {
	CreateProveedorUseCase create_proveedor.UseCase
}

func (handler CreateProveedor) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateProveedor) handle(ctx *gin.Context) {
	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	newProveedor, err := handler.CreateProveedorUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		case create_proveedor.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(newProveedor))

}

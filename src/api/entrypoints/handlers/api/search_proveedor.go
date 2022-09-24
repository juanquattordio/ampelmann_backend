package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_proveedor"
	"net/http"
	"strconv"
)

type SearchProveedor struct {
	SearchProveedorUseCase search_proveedor.UseCase
}

func (handler SearchProveedor) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler SearchProveedor) handle(ctx *gin.Context) {
	var request contracts.Request
	id := ctx.Query("id")
	if id != "" {
		proveedorId, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.Id = &proveedorId
	}
	cuit := ctx.Query("cuit")
	request.Cuit = &cuit

	err := request.Validate()
	if err != nil {
		ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		return
	}

	proveedorResult, err := handler.SearchProveedorUseCase.Execute(ctx, request.Id, request.Cuit)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, "Source not found"))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponse(proveedorResult))

}

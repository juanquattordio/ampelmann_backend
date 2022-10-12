package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_producto_final"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_producto_final"
	"net/http"
	"strconv"
)

type SearchProductoFinal struct {
	SearchProductoFinalUseCase search_producto_final.UseCase
}

func (handler SearchProductoFinal) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler SearchProductoFinal) handle(ctx *gin.Context) {
	var request contracts.Request

	id := ctx.Query("id_producto_final")
	if id != "" {
		productId, err := strconv.ParseInt(ctx.Query("id_producto_final"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.Id = &productId
	}
	descripcion := ctx.Query("descripcion")
	request.Descripcion = &descripcion

	err := request.Validate()
	if err != nil {
		ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		return
	}

	result, err := handler.SearchProductoFinalUseCase.Execute(ctx, request.Id, request.Descripcion)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, "Source not found"))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponse(result))

}

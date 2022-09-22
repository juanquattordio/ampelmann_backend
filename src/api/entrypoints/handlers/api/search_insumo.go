package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/search_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_insumo"
	"net/http"
	"strconv"
)

type SearchInsumo struct {
	SearchInsumoUseCase search_insumo.UseCase
}

func (handler SearchInsumo) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler SearchInsumo) handle(ctx *gin.Context) {
	var request contracts.Request

	id := ctx.Query("id_insumo")
	if id != "" {
		insumoId, err := strconv.ParseInt(ctx.Query("id_insumo"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.Id = &insumoId
	}
	nombre := ctx.Query("nombre")
	request.Nombre = &nombre

	err := request.Validate()
	if err != nil {
		ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		return
	}

	insumoResult, err := handler.SearchInsumoUseCase.Execute(ctx, request.Id, request.Nombre)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, "Source not found"))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponse(insumoResult))

}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_receta"
	"net/http"
	"strconv"
)

type UpdateReceta struct {
	UpdateRecetaUseCase update_receta.UseCase
}

func (handler UpdateReceta) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler UpdateReceta) handle(ctx *gin.Context) {
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

	recetaUpdated, err := handler.UpdateRecetaUseCase.Execute(ctx, id, request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
	}
	ctx.JSON(http.StatusOK, contracts.NewResponse(recetaUpdated))

}

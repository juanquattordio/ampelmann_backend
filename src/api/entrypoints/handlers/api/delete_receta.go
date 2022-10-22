package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/delete_receta"
	"net/http"
	"strconv"
)

type DeleteReceta struct {
	DeleteRecetaUseCase delete_receta.UseCase
}

func (handler DeleteReceta) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler DeleteReceta) handle(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, web.NewResponse(http.StatusBadRequest, nil, "Parametro incorrecto"))
		return
	}

	err = handler.DeleteRecetaUseCase.Execute(ctx, id)

	if err != nil {
		switch err {
		case delete_receta.ErrNotFound:
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("receta id %d not found", id)))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	} else {
		ctx.JSON(http.StatusOK, "success")
	}

}

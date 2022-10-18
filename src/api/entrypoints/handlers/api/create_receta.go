package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_receta"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_receta"
	"net/http"
)

type CreateReceta struct {
	CreateRecetaUseCase create_receta.UseCase
}

func (handler CreateReceta) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateReceta) handle(ctx *gin.Context) {
	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	receta, err := handler.CreateRecetaUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		// todo mejorar estos errores
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(receta))

}

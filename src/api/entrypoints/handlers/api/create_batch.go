package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_batch"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_batch"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/delete_receta"
	"net/http"
)

type CreateBatch struct {
	CreateBatchUseCase create_batch.UseCase
}

func (handler CreateBatch) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateBatch) handle(ctx *gin.Context) {
	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	batch, err := handler.CreateBatchUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		// todo mejorar estos errores
		case delete_receta.ErrNotFound:
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(batch))

}

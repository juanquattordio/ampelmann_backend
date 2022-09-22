package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_insumo"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_insumo"
	"net/http"
)

type CreateInsumo struct {
	CreateInsumoUseCase create_insumo.UseCase
}

func (handler CreateInsumo) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateInsumo) handle(ctx *gin.Context) {

	//token := ctx.Request.Header.Get("token")
	//fmt.Printf("token: " + token)
	//fmt.Printf("token getEnv: " + os.Getenv("TOKEN"))
	//if token != os.Getenv("TOKEN") {
	//	ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
	//	return
	//}
	// Esto es delegando la validación al ShoulBind

	var request contracts.Request

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}
	newInsumo, err := handler.CreateInsumoUseCase.Execute(ctx, request)
	if err != nil {
		switch err {
		case create_insumo.ErrDuplicate:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(newInsumo))

}

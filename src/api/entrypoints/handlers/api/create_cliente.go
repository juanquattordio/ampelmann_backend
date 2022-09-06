package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_cliente"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/create_cliente"
	"net/http"
)

type CreateCliente struct {
	CreateClienteUseCase create_cliente.UseCase
}

func (handler CreateCliente) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler CreateCliente) handle(ctx *gin.Context) {

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
	newCliente, err := handler.CreateClienteUseCase.Execute(ctx, request)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		//		ctx.JSON(404, web.NewResponse(404, nil, "Error al ejecutar Store"))
		return
	}
	ctx.JSON(http.StatusCreated, contracts.NewResponse(newCliente))

}

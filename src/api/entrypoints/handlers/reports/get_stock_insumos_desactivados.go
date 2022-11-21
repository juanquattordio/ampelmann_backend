package reports

import (
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/reports"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/reports/insumos_reports"
	"net/http"
	"os"
)

type InsumosReports struct {
	InsumosReportsUseCase insumos_reports.UseCase
}

func (handler InsumosReports) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler InsumosReports) handle(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	fmt.Printf("token: " + token)
	fmt.Printf("token getEnv: " + os.Getenv("TOKEN"))
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		return
	}

	result, err := handler.InsumosReportsUseCase.GetStockInsumosDesactivados()

	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponseStockInsumosDesactivados(result))
	return
}

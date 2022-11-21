package reports_handler

import (
	"database/sql"
	goErrors "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/reports"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/reports"
	"net/http"
	"os"
	"time"
)

type FacturacionReports struct {
	ReportsUseCase reports.UseCase
}

func (handler FacturacionReports) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler FacturacionReports) handle(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	fmt.Printf("token: " + token)
	fmt.Printf("token getEnv: " + os.Getenv("TOKEN"))
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "token inválido"))
		return
	}

	dateFromStr := ctx.Query("date-from")
	dateToStr := ctx.Query("date-to")

	dateFrom, err := time.Parse("2006-01-02", dateFromStr)
	dateTo, err := time.Parse("2006-01-02", dateToStr)
	if err != nil {
		ctx.JSON(404, web.NewResponse(404, nil, "Date invalid parameters"))
		return
	}

	result, err := handler.ReportsUseCase.GetFacturacionBetweenDates(dateFrom, dateTo)

	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponseFacturacionBetweenDates(result, dateFrom, dateTo))
	return
}

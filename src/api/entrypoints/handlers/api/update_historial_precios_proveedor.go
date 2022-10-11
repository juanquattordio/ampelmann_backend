package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/update_proveedor"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/update_historial_precios_proveedor"
	"net/http"
	"time"
)

type UpdateHistorialPreciosProveedor struct {
	UpdateHistorialPreciosProveedor update_historial_precios_proveedor.UseCase
}

func (handler UpdateHistorialPreciosProveedor) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler UpdateHistorialPreciosProveedor) handle(ctx *gin.Context) {
	var req contracts.RequestUpdateHistorialPrecio

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
		return
	}

	err := handler.UpdateHistorialPreciosProveedor.Execute(ctx, req.IdProveedor, req.IdInsumo,
		req.PrecioUnitario, time.Time(req.Fecha), req.Status)
	if err != nil {
		if goErrors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, "Source not found"))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, nil)

}

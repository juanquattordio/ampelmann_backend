package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/get_stock_producto"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/get_stock_producto"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_producto_final"
	"net/http"
	"strconv"
)

type GetStockProducto struct {
	GetStockProductoUseCase get_stock_producto.UseCase
	SearchProducto          search_producto_final.UseCase
}

func (handler GetStockProducto) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler GetStockProducto) handle(ctx *gin.Context) {
	var request contracts.Request

	// lectura de parametros y verificación
	idProducto := ctx.Query("id_producto")
	idDeposito := ctx.Query("id_deposito")
	if idProducto != "" {
		productoId, err := strconv.ParseInt(ctx.Query("id_producto"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.IdProducto = &productoId
	}
	if idDeposito != "" {
		depositoId, err := strconv.ParseInt(ctx.Query("id_deposito"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.IdDeposito = &depositoId
	}

	err := request.Validate()
	if err != nil {
		ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
		return
	}

	// caso de Stock por Producto
	if request.IdProducto != nil {
		producto, deposito, errors := handler.GetStockProductoUseCase.GetStockByProducto(ctx, request.IdProducto,
			request.IdDeposito)
		if errors != nil {
			if goErrors.Is(errors, sql.ErrNoRows) {
				ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
			} else {
				ctx.JSON(500, web.NewResponse(500, nil, errors.Error()))
			}
			return
		}
		ctx.JSON(http.StatusOK, contracts.NewResponse(producto, deposito))
		return
	}

	// caso de Stock por depósito
	deposito, productos, errors := handler.GetStockProductoUseCase.GetStockByDeposito(ctx, request.IdDeposito)
	if errors != nil {
		if goErrors.Is(errors, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponseByDeposito(deposito, productos))
	return

}

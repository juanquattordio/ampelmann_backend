package api

import (
	"database/sql"
	goErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/juanquattordio/ampelmann_backend/src/api/config/web"
	contracts "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/get_stock"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/get_stock"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/search_insumo"
	"net/http"
	"strconv"
)

type GetStockInsumo struct {
	GetStockUseCase get_stock.UseCase
	SearchInsumo    search_insumo.UseCase
}

func (handler GetStockInsumo) Handle(ginContext *gin.Context) {
	handler.handle(ginContext)
}

func (handler GetStockInsumo) handle(ctx *gin.Context) {
	var request contracts.Request

	// lectura de parametros y verificación
	idInsumo := ctx.Query("id_insumo")
	idDeposito := ctx.Query("id_deposito")
	if idInsumo != "" {
		insumoId, err := strconv.ParseInt(ctx.Query("id_insumo"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		request.IdInsumo = &insumoId
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

	// caso de Stock por Insumo
	if request.IdInsumo != nil {
		insumo, deposito, errors := handler.GetStockUseCase.GetStockByInsumo(ctx, request.IdInsumo, request.IdDeposito)
		if errors != nil {
			if goErrors.Is(errors, sql.ErrNoRows) {
				ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
			} else {
				ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
			}
			return
		}
		ctx.JSON(http.StatusOK, contracts.NewResponse(insumo, deposito))
		return
	}

	// caso de Stock por depósito
	deposito, insumos, errors := handler.GetStockUseCase.GetStockByDeposito(ctx, request.IdDeposito)
	if errors != nil {
		if goErrors.Is(errors, sql.ErrNoRows) {
			ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
		} else {
			ctx.JSON(404, web.NewResponse(404, nil, errors.Error()))
		}
		return
	}
	ctx.JSON(http.StatusOK, contracts.NewResponseByDeposito(deposito, insumos))
	return

}

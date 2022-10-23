package create_batch

import (
	"context"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_batch"
	movimientoDeposito "github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/movimiento_depositos"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/errors/apierrors"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/usecases/movimiento_depositos"
	"time"
)

type Implementation struct {
	BatchProvider            providers.Batch
	RecetaProvider           providers.Receta
	MovimientoInsumosUseCase movimiento_depositos.UseCase
}

func (uc *Implementation) Execute(ctx context.Context, request create_batch.Request) (*entities.Batch, error) {
	if *request.LitrosProducidos == 0 {
		return nil, errors.NewBadRequest(apierrors.BadRequestMessage)
	}
	// verifica que exista la receta y calcula los insumos necesarios para los litros finales requeridos
	ingredientes, err := uc.RecetaProvider.CalculateIngredientes(request.IdReceta, *request.LitrosProducidos)
	if err != nil {
		return nil, err
	}

	// verifica que haya stock suficiente para la producción y carga el movimiento de insumos
	var reqMov movimientoDeposito.Request
	if err = reqConstructor(uc, &reqMov, ingredientes); err != nil {
		return nil, err
	}
	_, err = uc.MovimientoInsumosUseCase.Execute(ctx, reqMov)
	if err != nil {
		return nil, err
	}

	fecha := time.Now()
	if request.FechaOrigen != nil {
		fecha = time.Time(*request.FechaOrigen)
	}

	newBatch := entities.Batch{IdReceta: *request.IdReceta, Fecha: fecha, LitrosProducidos: *request.LitrosProducidos}

	if err = uc.BatchProvider.CreateBatch(&newBatch); err != nil {
		return nil, goErrors.New(fmt.Sprintf("fallo en la creación del batch"))
	}

	//// Crea un movimiento que ejecuta Updates de stocks en cada depósito
	//idDepositoInsumos := int64(2)
	//causaMovimiento := fmt.Sprintf("FCP-%d", newBatch.IdBatch)
	//movimiento := entities.NewMovimientoDeposito(0, idDepositoInsumos, parseToMovLines(request.Lineas), causaMovimiento)
	//if err = uc.StockProvider.MovimientoDepositos(ctx, movimiento); err != nil {
	//	return nil, goErrors.New(fmt.Sprintf("fallo en la creación del movimiento de insumos por compra"))
	//}

	return &newBatch, nil
}

func reqConstructor(uc *Implementation, req *movimientoDeposito.Request, ingredientes []entities.Ingredientes) error {
	req.IdDepositoOrigen = int64(2)  // Insumos
	req.IdDepositoDestino = int64(0) // A descontar
	insumos := make([]movimientoDeposito.Insumo, len(ingredientes))
	for i := range ingredientes {
		insumos[i].IdLinea = int64(i + 1)
		insumos[i].IdInsumo = &ingredientes[i].IdInsumo
		insumos[i].Cantidad = &ingredientes[i].Cantidad
		insumos[i].Observaciones = ingredientes[i].Observaciones
	}
	req.Insumos = insumos

	lastBatch, err := uc.BatchProvider.GetLastBacth()
	if err != nil {
		return err
	}

	req.CausaMovimiento = fmt.Sprintf("OP-%d", lastBatch+1)
	return nil
}

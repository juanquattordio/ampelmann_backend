package create_batch

import (
	"context"
	goErrors "errors"
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/contracts/create_batch"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/providers"
	"time"
)

type Implementation struct {
	BatchProvider  providers.Batch
	RecetaProvider providers.Receta
}

func (uc *Implementation) Execute(ctx context.Context, request create_batch.Request) (*entities.Batch, error) {
	// valida que exista la receta
	_, err := uc.RecetaProvider.Search(request.IdReceta)
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

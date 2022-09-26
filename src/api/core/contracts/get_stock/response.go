package get_stock

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ResponseByDeposito struct {
	ID          string        `json:"deposito_id"`
	Descripcion string        `json:"descripcion"`
	Insumos     []InsumoStock `json:"insumos"`
}

type Response struct {
	Deposito   string  `json:"deposito,omitempty"`
	ID         string  `json:"insumo_id"`
	Nombre     string  `json:"insumo_nombre"`
	StockTotal float64 `json:"total_stock"`
}
type InsumoStock struct {
	ID     string  `json:"insumo_id"`
	Nombre string  `json:"insumo_nombre"`
	Stock  float64 `json:"stock"`
}

func NewResponse(insumo *entities.Insumo, deposito *entities.Deposito) *Response {
	return &Response{
		Deposito:   deposito.Descripcion,
		ID:         fmt.Sprintf("%d", insumo.IdInsumo),
		Nombre:     insumo.Nombre,
		StockTotal: insumo.Stock,
	}
}
func NewResponseByDeposito(deposito *entities.Deposito, insumos []entities.Insumo) *ResponseByDeposito {
	insumosList := make([]InsumoStock, len(insumos))
	for i := range insumos {
		insumosList[i].ID = fmt.Sprintf("%d", insumos[i].IdInsumo)
		insumosList[i].Nombre = insumos[i].Nombre
		insumosList[i].Stock = insumos[i].Stock
	}
	return &ResponseByDeposito{
		ID:          fmt.Sprintf("%d", deposito.ID),
		Descripcion: deposito.Descripcion,
		Insumos:     insumosList,
	}
}

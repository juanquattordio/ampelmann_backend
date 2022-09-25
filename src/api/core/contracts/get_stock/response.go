package get_stock

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ResponseByInsumo struct {
	ID            string          `json:"insumo_id"`
	Nombre        string          `json:"insumo_nombre"`
	DepositoStock []DepositoStock `json:"deposito_stock"`
}
type DepositoStock struct {
	ID          string  `json:"deposito_id"`
	Descripcion string  `json:"descripcion"`
	Stock       float64 `json:"stock"`
}

type ResponseByDeposito struct {
	ID          string        `json:"deposito_id"`
	Descripcion string        `json:"descripcion"`
	InsumoStock []InsumoStock `json:"insumo_stock"`
}
type InsumoStock struct {
	ID     string  `json:"insumo_id"`
	Nombre string  `json:"insumo_nombre"`
	Stock  float64 `json:"stock"`
}

type Response struct {
	Deposito   string  `json:"deposito,omitempty"`
	ID         string  `json:"insumo_id"`
	Nombre     string  `json:"insumo_nombre"`
	StockTotal float64 `json:"total_stock"`
}

func NewResponseByInsumo(insumo *entities.Insumo) *ResponseByInsumo {
	return &ResponseByInsumo{
		ID:     fmt.Sprintf("%d", insumo.IdInsumo),
		Nombre: insumo.Nombre,
		//TODO modificar esto
		DepositoStock: nil,
	}
}
func NewResponse(insumo *entities.Insumo, deposito *entities.Deposito, stock float64) *Response {
	return &Response{
		Deposito:   deposito.Descripcion,
		ID:         fmt.Sprintf("%d", insumo.IdInsumo),
		Nombre:     insumo.Nombre,
		StockTotal: stock,
	}
}

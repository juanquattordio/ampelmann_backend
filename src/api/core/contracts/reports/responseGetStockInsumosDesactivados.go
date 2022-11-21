package reports

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ResponseInsumosDesactivados struct {
	InsumosDesactivados []InsumoDesactivado `json:"insumos_desactivados"`
}

type InsumoDesactivado struct {
	IdInsumo string  `json:"insumo_id"`
	Nombre   string  `json:"insumo_nombre"`
	Stock    float64 `json:"stock"`
	Status   string  `json:"status"`
}

func NewResponseStockInsumosDesactivados(insumos []entities.Insumo) *ResponseInsumosDesactivados {
	insumosList := make([]InsumoDesactivado, len(insumos))
	for i := range insumos {
		insumosList[i].IdInsumo = fmt.Sprintf("%d", insumos[i].IdInsumo)
		insumosList[i].Nombre = insumos[i].Nombre
		insumosList[i].Stock = insumos[i].Stock
		insumosList[i].Status = insumos[i].Status
	}
	return &ResponseInsumosDesactivados{
		InsumosDesactivados: insumosList,
	}
}

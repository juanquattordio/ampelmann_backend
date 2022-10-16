package get_stock_producto

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ResponseByDeposito struct {
	ID          string          `json:"deposito_id"`
	Descripcion string          `json:"descripcion"`
	Productos   []ProductoStock `json:"productos"`
}

type Response struct {
	IdDeposito  int64   `json:"id_deposito,omitempty"`
	Deposito    string  `json:"deposito,omitempty"`
	ID          string  `json:"id_producto"`
	Descripcion string  `json:"producto_descripcion"`
	StockTotal  float64 `json:"total_stock"`
}
type ProductoStock struct {
	ID          string  `json:"id_producto"`
	Descripcion string  `json:"producto_descripcion"`
	Stock       float64 `json:"stock"`
}

func NewResponse(producto *entities.ProductoFinal, deposito *entities.Deposito) *Response {
	//if deposito.ID==0
	return &Response{
		IdDeposito:  deposito.ID,
		Deposito:    deposito.Descripcion,
		ID:          fmt.Sprintf("%d", producto.Id),
		Descripcion: producto.Descripcion,
		StockTotal:  producto.Stock,
	}
}
func NewResponseByDeposito(deposito *entities.Deposito, productos []entities.ProductoFinal) *ResponseByDeposito {
	productosList := make([]ProductoStock, len(productos))
	for i := range productos {
		productosList[i].ID = fmt.Sprintf("%d", productos[i].Id)
		productosList[i].Descripcion = productos[i].Descripcion
		productosList[i].Stock = productos[i].Stock
	}
	return &ResponseByDeposito{
		ID:          fmt.Sprintf("%d", deposito.ID),
		Descripcion: deposito.Descripcion,
		Productos:   productosList,
	}
}

package reports

import (
	"fmt"
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

type ResponseProductosDesactivados struct {
	ProductosDesactivados []ProductoDesactivado `json:"productos_desactivados"`
}

type ProductoDesactivado struct {
	IdProducto string  `json:"producto_id"`
	Nombre     string  `json:"descripcion"`
	Stock      float64 `json:"stock"`
	Status     string  `json:"status"`
}

func NewResponseStockProductosDesactivados(productos []entities.ProductoFinal) *ResponseProductosDesactivados {
	productosList := make([]ProductoDesactivado, len(productos))
	for i := range productos {
		productosList[i].IdProducto = fmt.Sprintf("%d", productos[i].Id)
		productosList[i].Nombre = productos[i].Descripcion
		productosList[i].Stock = productos[i].Stock
		productosList[i].Status = productos[i].Status
	}
	return &ResponseProductosDesactivados{
		ProductosDesactivados: productosList,
	}
}

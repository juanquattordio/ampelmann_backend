package reports_repository

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
	"github.com/juanquattordio/ampelmann_backend/src/api/repositories/cliente"
)

const (
	getStockInsumosDesactivados     = "SELECT I.idInsumo as 'idArticulo', I.nombre, I.unidad_medida, SUM(SI.stock) as 'total_stock', I.status FROM Insumo I INNER JOIN Stock_Insumo SI on I.idInsumo = SI.id_insumo GROUP BY I.idInsumo, I.nombre, I.status HAVING I.status = 'desactivo' AND total_stock>0"
	getStockProductosDesactivados   = "SELECT P.id_producto as 'idArticulo', P.descripcion as 'nombre', P.unidad_medida, SUM(SP.stock) as 'total_stock', P.status FROM Producto_Final P INNER JOIN Stock_Producto SP on P.id_producto = SP.id_producto GROUP BY P.id_producto, P.descripcion, P.status HAVING P.status = 'desactivo' AND total_stock>0"
	getClientesDesactivados         = "SELECT idCliente, cuit, nombre, ubicacion, email, status FROM Cliente WHERE status = 'desactivo'"
	getFacturacionTotalBetweenDates = "SELECT SUM(importe_total) as 'importe_total' FROM Venta_Factura_Header WHERE fecha BETWEEN ? AND ?"
)

type articulo struct {
	ID     int64   `db:"idArticulo"`
	Nombre string  `db:"nombre"`
	Unidad string  `db:"unidad_medida"`
	Stock  float64 `db:"total_stock"`
	Status string  `db:"status"`
}

func (dbItem articulo) toEntityInsumo() *entities.Insumo {
	return &entities.Insumo{
		IdInsumo: dbItem.ID,
		Nombre:   dbItem.Nombre,
		Unidad:   dbItem.Unidad,
		Stock:    dbItem.Stock,
		Status:   dbItem.Status,
	}
}

func toEntitiesInsumos(insumos []articulo) []entities.Insumo {
	var insumosResult []entities.Insumo
	for _, insumoDb := range insumos {
		i := new(entities.Insumo)
		i.IdInsumo = insumoDb.ID
		i.Nombre = insumoDb.Nombre
		i.Unidad = insumoDb.Unidad
		i.Stock = insumoDb.Stock
		i.Status = insumoDb.Status
		insumosResult = append(insumosResult, *i)
	}
	return insumosResult
}

func toEntitiesProductos(articulos []articulo) []entities.ProductoFinal {
	var result []entities.ProductoFinal
	for _, articuloDb := range articulos {
		i := new(entities.ProductoFinal)
		i.Id = articuloDb.ID
		i.Descripcion = articuloDb.Nombre
		i.Unidad = articuloDb.Unidad
		i.Stock = articuloDb.Stock
		i.Status = articuloDb.Status
		result = append(result, *i)
	}
	return result
}

func toEntitiesClientes(clientes []cliente.Cliente) []entities.Cliente {
	var clientesResult []entities.Cliente
	for _, clienteDb := range clientes {
		c := new(entities.Cliente)
		c.ID = clienteDb.ID
		c.Nombre = clienteDb.Nombre
		c.Cuit = clienteDb.Cuit
		c.Ubicacion = clienteDb.Ubicacion
		c.Email = clienteDb.Email
		c.Status = clienteDb.Status
		clientesResult = append(clientesResult, *c)
	}
	return clientesResult
}

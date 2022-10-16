package stock_producto

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "stockProducto"
)

const (
	getStockProductoDeposito = "SELECT id_deposito, id_producto, stock FROM Stock_Producto"
	sumStockByProducto       = "SELECT id_producto, SUM(stock) as 'stock' from Stock_Producto GROUP BY id_producto HAVING id_producto = ?"
	getStockByDeposito       = "select sp.id_deposito, pf.id_producto, pf.descripcion as 'descripcion', sp.stock as 'stock' from Stock_Producto sp INNER JOIN Producto_Final pf ON sp.id_producto = pf.id_producto WHERE sp.id_deposito= ?"
	updateStock              = `INSERT INTO Stock_Producto (id_deposito, id_insumo, stock) VALUES (:id_deposito,:id_insumo,:stock+stock) ON DUPLICATE KEY UPDATE stock = VALUES(stock)+:stock`
)

type stockProducto struct {
	IdDeposito int64   `db:"id_deposito"`
	IdProducto int64   `db:"id_producto"`
	Stock      float64 `db:"stock"`
}

func (dbItem stockProducto) toEntity() *ResponseStock {
	return &ResponseStock{
		IdDeposito: dbItem.IdDeposito,
		IdProducto: dbItem.IdProducto,
		Stock:      dbItem.Stock,
	}
}

func toDBEntity(idDeposito int64, idProducto int64, cantidad float64) *stockProducto {
	return &stockProducto{
		IdDeposito: idDeposito,
		IdProducto: idProducto,
		Stock:      cantidad,
	}
}

type ResponseStock struct {
	IdDeposito int64
	IdProducto int64
	Stock      float64
}

type stockDeposito struct {
	IdDeposito  int64   `db:"id_deposito"`
	IdProducto  int64   `db:"id_producto"`
	Descripcion string  `db:"descripcion"`
	Stock       float64 `db:"stock"`
}

func (dbItem stockDeposito) toEntity() entities.ProductoFinal {
	return entities.ProductoFinal{
		Id:          dbItem.IdProducto,
		Descripcion: dbItem.Descripcion,
		Stock:       dbItem.Stock,
	}
}

type ResponseStockDeposito struct {
	IdDeposito  int64
	IdProducto  int64
	Descripcion string
	Stock       float64
}

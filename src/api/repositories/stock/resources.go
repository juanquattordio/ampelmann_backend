package stock

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "stockInsumo"
)

const (
	getStockInsumoDeposito = "SELECT id_insumo, id_deposito, stock FROM Stock_Insumo"
	sumStockByInsumo       = "SELECT id_insumo, SUM(stock) as 'stock' from Stock_Insumo GROUP BY id_insumo HAVING id_insumo = ?"
	getStockByDeposito     = "select id_deposito, id_insumo, i.nombre as 'nombre', Stock_Insumo.stock as 'stock' from Stock_Insumo INNER JOIN Insumo i ON id_insumo = i.idInsumo WHERE id_deposito= ?"
	updateStockInsumos     = "call updateStockInsumos(?,?,?)"   // llama al procedure/procedimiento creado en la BD
	updateStockProductos   = "call updateStockProductos(?,?,?)" // llama al procedure/procedimiento creado en la BD
)

type stockInsumo struct {
	IdDeposito int64   `db:"id_deposito"`
	IdInsumo   int64   `db:"id_insumo"`
	Stock      float64 `db:"stock"`
}

func (dbItem stockInsumo) toEntity() *ResponseStock {
	return &ResponseStock{
		IdDeposito: dbItem.IdDeposito,
		IdInsumo:   dbItem.IdInsumo,
		Stock:      dbItem.Stock,
	}
}

type ResponseStock struct {
	IdDeposito int64
	IdInsumo   int64
	Stock      float64
}

type stockDeposito struct {
	IdDeposito   int64   `db:"id_deposito"`
	IdInsumo     int64   `db:"id_insumo"`
	NombreInsumo string  `db:"nombre"`
	Stock        float64 `db:"stock"`
}

func (dbItem stockDeposito) toEntity() entities.Insumo {
	return entities.Insumo{
		IdInsumo: dbItem.IdInsumo,
		Nombre:   dbItem.NombreInsumo,
		Stock:    dbItem.Stock,
	}
}

type ResponseStockDeposito struct {
	IdDeposito   int64
	IdInsumo     int64
	NombreInsumo string
	Stock        float64
}

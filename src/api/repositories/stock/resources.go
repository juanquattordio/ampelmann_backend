package stock

const (
	resource = "stock"
)

const (
	selectScriptMySQL = "SELECT id_insumo, id_deposito, stock FROM Stock_Insumo"
	sumStockByInsumo  = "SELECT id_insumo, SUM(stock) as 'stock' from Stock_Insumo GROUP BY id_insumo HAVING id_insumo = ?"
)

type stock struct {
	IdDeposito int64   `db:"id_deposito"`
	IdInsumo   int64   `db:"id_insumo"`
	Stock      float64 `db:"stock"`
}

func (dbItem stock) toEntity() *ResponseStock {
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

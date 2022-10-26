package documento

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "stockInsumo"
)

const (
	insertMovInsumoHeader     = "INSERT INTO Movimiento_Insumos_Header (idDeposito_Origen, idDeposito_Destino, fecha, causa_movimiento) VALUES(?,?,?,?)"
	insertMovProductoHeader   = "INSERT INTO Movimiento_Productos_Header (id_deposito_origen, id_deposito_destino, fecha, causa_movimiento) VALUES(?,?,?,?)"
	insertMovInsumoLine       = "INSERT INTO Movimiento_Insumos_Line (idMovimiento, idLinea, idInsumo, cantidad, observaciones) VALUES(?,?,?,?,?)"
	insertMovProductoLine     = "INSERT INTO Movimiento_Productos_Line (id_movimiento, id_linea, id_producto, cantidad, observaciones) VALUES(?,?,?,?,?)"
	insertFacturaCompraHeader = "INSERT INTO Compra_Factura_Header (id_proveedor, id_factura_proveedor, fecha_origen, fecha, importe_total) VALUES(?,?,?,?,?)"
	insertFacturaCompraLine   = "INSERT INTO Compra_Factura_Line (id_factura, id_linea, id_insumo, cantidad, precio_unitario, observaciones) VALUES(?,?,?,?,?,?)"
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

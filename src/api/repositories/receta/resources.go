package receta

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "stockInsumo"
)

const (
	insertRecetaHeader = "INSERT INTO Receta_Header (descripcion_paso_paso, id_producto, litros_finales) VALUES(?,?,?)"
	insertRecetaLine   = "INSERT INTO Receta_Insumo (id_receta, id_insumo, cantidad_insumo, observaciones) VALUES(?,?,?,?)"
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

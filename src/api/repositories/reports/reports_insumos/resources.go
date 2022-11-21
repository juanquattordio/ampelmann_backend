package reports_insumos

import (
	"github.com/juanquattordio/ampelmann_backend/src/api/core/entities"
)

const (
	getStockInsumosDesactivados = "SELECT I.idInsumo, I.nombre, I.unidad_medida, SUM(SI.stock) as 'total_stock', I.status FROM Insumo I INNER JOIN Stock_Insumo SI on I.idInsumo = SI.id_insumo GROUP BY I.idInsumo, I.nombre, I.status HAVING I.status = 'desactivo' AND total_stock>0"
)

type insumo struct {
	ID     int64   `db:"idInsumo"`
	Nombre string  `db:"nombre"`
	Unidad string  `db:"unidad_medida"`
	Stock  float64 `db:"total_stock"`
	Status string  `db:"status"`
}

func (dbItem insumo) toEntity() *entities.Insumo {
	return &entities.Insumo{
		IdInsumo: dbItem.ID,
		Nombre:   dbItem.Nombre,
		Unidad:   dbItem.Unidad,
		Stock:    dbItem.Stock,
		Status:   dbItem.Status,
	}
}

func toEntities(insumos []insumo) []entities.Insumo {
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

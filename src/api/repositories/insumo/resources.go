package insumo

import "github.com/juanquattordio/ampelmann_backend/src/api/core/entities"

const (
	resource = "insumo"
)

const (
	saveScriptMySQL = "INSERT INTO Insumo(nombre, stock, status)" +
		"VALUES(?, ?, ?) "
	selectScriptMySQL = "SELECT idInsumo, nombre, stock, status FROM Insumo"
	updateScriptMySQL = "UPDATE Insumo SET nombre = ?, stock = ?, status = ? WHERE idInsumo = ?"
)

type insumo struct {
	ID     int64   `db:"idInsumo"`
	Nombre string  `db:"nombre"`
	Stock  float64 `db:"stock"`
	Status string  `db:"status"`
}

func newEntity(i entities.Insumo) insumo {
	return insumo{
		ID:     i.IdInsumo,
		Nombre: i.Nombre,
		Stock:  i.Stock,
		Status: i.Status,
	}
}

func (dbItem insumo) toEntity() *entities.Insumo {
	return &entities.Insumo{
		IdInsumo: dbItem.ID,
		Nombre:   dbItem.Nombre,
		Stock:    dbItem.Stock,
		Status:   dbItem.Status,
	}
}
